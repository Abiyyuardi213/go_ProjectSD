document.addEventListener("DOMContentLoaded", function() {
    document.getElementById("login-form").addEventListener("submit", function(event) {
        event.preventDefault(); // Mencegah form submit default

        var username = document.getElementById('login-username').value;
        var password = document.getElementById('login-password').value;

        var xhr = new XMLHttpRequest();
        xhr.open("POST", "/login", true);
        xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4) {
                if (xhr.status === 200) {
                    alert("Login successful: " + xhr.responseText);
                } else {
                    alert("Error: " + xhr.responseText);
                }
            }
        };
        xhr.send(JSON.stringify({ username: username, password: password }));
    });
});
