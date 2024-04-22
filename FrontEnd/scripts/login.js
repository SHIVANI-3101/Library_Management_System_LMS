document.getElementById('loginform').addEventListener('submit', function(e) {
    e.preventDefault();

    const url = `http://localhost:8080/user/login`;

    const data = {
        "email": document.getElementById('email').value,
        "password": document.getElementById('password').value,
    };

    const requestOptions = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    };

    fetch(url, requestOptions)
        .then(response => {
            if (!response.ok) {
                throw new Error('Login failed');
            }
            return response.json();
        })
        .then(data => {
            if (data.message === 'Login successful' && data.token) {
                // Set the token as a cookie
                setCookie('token', data.token, 1); // Expires in 1 day
                alert('Login successful');

                // Redirect based on user role
                if(data.role == 'Creator') {
                    window.location.href = 'librarylisting.html';
                } else if(data.role == 'Admin') {
                    window.location.href = 'booklisting.html';
                } else if(data.role == 'Reader') {
                    window.location.href = 'bookrequestlisting.html';
                }
            } else {
                alert(data.error || 'Login failed');
            }
        })
        .catch(error => {
            console.error('Error fetching data:', error);
            alert('Login failed');
        });
});

// Function to set a cookie
function setCookie(name, value, days) {
    var expires = "";
    if (days) {
        var date = new Date();
        date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
        expires = "; expires=" + date.toUTCString();
    }
    document.cookie = name + "=" + (value || "") + expires + "; path=/";
}
