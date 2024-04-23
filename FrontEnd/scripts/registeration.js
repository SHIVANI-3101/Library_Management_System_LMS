document.getElementById('registerationform').addEventListener('submit',function(e)
{
    e.preventDefault();
    alert('This is a cat');

    const url = `http://localhost:8080/user/create`;

    data = {
        "name": document.getElementById('name').value,
        "email": document.getElementById('email').value,
        "contact_number": parseInt(document.getElementById('contact').value),
        "role": document.getElementById('rolevalue').value,
        "pass": document.getElementById('password').value
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
            if(data.message=='success')
            {
                alert('Thankyou for registering with us');
                window.location.href = 'login.html';
            }
        })
        .catch(error => console.error('Error fetching data:', error));
})