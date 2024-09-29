
console.log("JS is working")

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
