(function(){
    const buttonThemeToogle = document.querySelector('#theme-toggle');
    buttonThemeToogle.addEventListener('click', () => {
        document.documentElement.classList.toggle(
            "dark"
        );
        document.documentElement.classList.contains("dark")
        localStorage.theme = document.documentElement.classList.contains("dark") ? "dark" : "light";
    })
}())