document.getElementById('loginform').addEventListener('submit', function(e) {
    e.preventDefault();

    const url = `http://localhost:8080/user/login`;

    data = {
        "email": document.getElementById('email').value,
        "password": document.getElementById('password').value,
    }

    const requestOptions = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    };

    fetch(url, requestOptions)
        .then(response => response.json())
        .then(data => {
            if (data.error) {
                alert(data.error); 
            } 
            else if (data.message === 'Login successful') 
            {
                alert('Login successful');

                if(data.role=='Creator')
                {
                    window.location.href = 'librarylisting.html';
                }
                else if(data.role=='Admin')
                {
                    window.location.href = 'booklisting.html';
                }
                else if(data.role=='Reader')
                {
                    window.location.href = 'bookrequestlisting.html';
                }
            }
        })
        .catch(error => console.error('Error fetching data:', error));
});
