
console.log("JS is working")

// Firebase
import { initializeApp } from 'https://www.gstatic.com/firebasejs/11.3.1/firebase-app.js';
import { getAuth, signOut, GoogleAuthProvider, signInWithPopup } from 'https://www.gstatic.com/firebasejs/11.3.1/firebase-auth.js';

const firebaseConfig = {
    apiKey: "AIzaSyBZjxfmLXlYfnVraD55jw1XBOTMbp04ZbQ",
    authDomain: "thefatbean-c667b.firebaseapp.com",
    projectId: "thefatbean-c667b",
    storageBucket: "thefatbean-c667b.firebasestorage.app",
    messagingSenderId: "899089413678",
    appId: "1:899089413678:web:224ac7bbe0685e1bf391bf",
    measurementId: "G-SF9V6KKSNX"
};

const app = initializeApp(firebaseConfig);
const auth = getAuth(app);
document.getElementById('sign-out-btn')?.addEventListener('click', function (event) {
    signOut(auth).then(() => {
        window.location.href = "/auth/sign-out";
    }).catch((error) => {
        console.error("Error signing out:", error);
    });
});

// Google Sign-in
const provider = new GoogleAuthProvider();
document.getElementById('google-auth-btn')?.addEventListener('click', function() {
    signInWithPopup(auth, provider)
        .then((result) => {
            const user = result.user;
            console.log('User signed in:', user);
            user.getIdToken().then((idToken) => {
                console.log('Token :', idToken);
                const redirectUrl = `/auth/google-callback?idToken=${encodeURIComponent(idToken)}`;
                window.location.href = redirectUrl;
            });
        })
        .catch((error) => {
            console.error('Sign-in error:', error);
        });
});


// Main
const toggleDarkModeButton = document.getElementById('dark-mode-toggle');
const body = document.body;

toggleDarkModeButton.addEventListener('click', () => {
    body.classList.toggle('dark-mode');
    const theme = body.classList.contains('dark-mode') ? 'dark' : 'light';
    localStorage.setItem('theme', theme);
    console.log("Dark Mode " + theme);
});

// Check user preference and set theme on page load
const theme = localStorage.getItem('theme');
if (theme === 'dark') {
    body.classList.add('dark-mode');
    toggleDarkModeButton.checked = true;
}
