const express = require('express');
const app = express();
const path = require('path');
const PORT = process.env.PORT || 3000;
const config = require('./config');
const session = require('express-session');
const fs = require('fs');


app.use(session({
    secret: config.SESSION_SECRET,
    resave: false,
    saveUninitialized: true,
    cookie: { secure: false }
}));

app.get('/', (req, res) => {
    if (req.session.accessToken) {
        return res.redirect('/feed');
    }
    res.sendFile(path.join(__dirname, '../client/index.html'));
});

app.get('/login', (req, res) => {
    return res.redirect(config.AUTH_URL);
});

app.get('/auth', async (req, res) => {
    const code = req.query.code;
    try {
        const tokenResponse = await fetch('https://api.intra.42.fr/oauth/token', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                grant_type: 'authorization_code',
                client_id: config.CLIENT_ID,
                client_secret: config.CLIENT_SECRET,
                redirect_uri: config.REDIRECT_URI,
                code: code,
            }),
        });
    
        if (!tokenResponse.ok) {
            throw new Error('Failed to fetch access token');
        }
    
        const tokenData = await tokenResponse.json();
        const accessToken = tokenData.access_token;
    
        if (!accessToken) {
            throw new Error('No access token');
        }
    
        const meResponse = await fetch('https://api.intra.42.fr/v2/me', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${accessToken}`,
            },
        });
    
        if (!meResponse.ok) {
            throw new Error('Invalid token');
        }
    
        const meData = await meResponse.json();
    
        req.session.accessToken = accessToken;
        res.redirect('/feed');
    } catch (error) {
        console.error(error);
        res.status(500).send('An error occurred');
    }
    
});

app.get('/feed', (req, res) => {
    if (!req.session.accessToken) {
        return res.redirect('/');
    }
        res.sendFile(path.join(__dirname, '../client/routes/feed.html'));
});

app.get('/friends', (req, res) => {
    if (!req.session.accessToken) {
        return res.redirect('/');
    }
        res.sendFile(path.join(__dirname, '../client/routes/friends.html'));
});

app.get('/friends/add', (req, res) => {
    if (!req.session.accessToken) {
        return res.redirect('/');
    }
        res.sendFile(path.join(__dirname, '../client/routes/add.html'));
});

app.get('/logout', (req, res) => {
    req.session.destroy((err) => {
        if (err) {
            return res.status(500).send('Failed to logout');
        }
        res.redirect('/');
    });
});

app.use(express.static(path.join(__dirname, '../client')));

app.listen(PORT, () => {
    console.log("Server is running on port " + PORT);
});