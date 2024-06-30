
console.log("JS is working")

const toggleDarkModeButton = document.getElementById('toggle-dark-mode');
const body = document.body;

toggleDarkModeButton.addEventListener('click', () => {
    body.classList.toggle('dark-mode');
    const theme = body.classList.contains('dark-mode') ? 'dark' : 'light';
    if (theme === "dark"){
        toggleDarkModeButton.innerHTML = "Light Mode";
    }
    else{
        toggleDarkModeButton.innerHTML = "Dark Mode";
    }
    localStorage.setItem('theme', theme);
});

// Check user preference and set theme on page load
const theme = localStorage.getItem('theme');
if (theme === 'dark') {
    body.classList.add('dark-mode');
    toggleDarkModeButton.innerHTML = "Light Mode";
}
