let libraryID = null;
let action = null;
const token = getCookie('token');

function initialSetup()
{
    const urlParams = new URLSearchParams(window.location.search);
    libraryID = urlParams.get('id');
    action = urlParams.get('action');

    if(action=='create')
    {
        // Let the cotrol pass
        document.getElementById('librarycreationform').addEventListener('submit',function(e)
        {
            e.preventDefault();
            libraryAction();
        })
    }
    else if(action=='edit')
    {
        document.getElementById('librarycreationform').addEventListener('submit',function(e)
        {
            e.preventDefault();
            libraryAction();
        })
        getSpecificlibraryData();
    }
    else if(action==null)
    {
        getAllData();
        document.getElementById('searchform').addEventListener('submit',function(e)
        {
            e.preventDefault();
            Search();
        })
        document.getElementById('resetButton').addEventListener('click',function(e)
        {
            location.reload();
        })
    }
}

function getAllData() {

    const url = 'http://localhost:8080/owner/libraries';
    const requestOptions = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}` 
        }
    };

    fetch(url, requestOptions)
        .then(response => response.json())
        .then(data => {
            const tableBody = document.querySelector('tbody');
            tableBody.innerHTML = ''; // Clear existing table rows
            
            if (data.length === 0) {
                const row = document.createElement('tr');
                row.innerHTML = `<td colspan="8">No libraries present</td>`;
                tableBody.appendChild(row);
            } else {
                let serialNo = 1;
                data.forEach(item => {
                    const row = document.createElement('tr');
                    
                    // Populate table cells
                    row.innerHTML = `
                        <td>${serialNo}</td>
                        <td>${item.name}</td>
                        <td>
                            <a class="btn btn-primary" href="createlibrary.html?action=edit&id=${item.id}">Edit</a>
                            <a class="btn btn-danger" href="#" data-library-id="${item.id}" onclick="deletelibrary(event)">Delete</a>
                        </td>
                    `;
                    
                    tableBody.appendChild(row);
                    serialNo++;
                });
            }
        })
        .catch(error => console.error('Error fetching data:', error));
}

function getSpecificlibraryData()
{
    const url = `http://localhost:8080/owner/library/${libraryID}`;

    document.getElementById('libraryformtitle').innerText = "Update library";
    document.getElementById('libraryformbutton').innerText = "Update";

    const requestOptions = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        }
    };

    fetch(url,requestOptions)
        .then(response => response.json())
        .then(data => {
            document.getElementById('library-id').value = data.id;
            document.getElementById('library-name').value = data.name;
            fetchAdministratorData(data.id);
        })
        .catch(error => console.error('Error fetching data:', error));
}

function fetchAdministratorData(libraryID)
{
    const url = `http://localhost:8080/owner/library/admin/${libraryID}`;

    document.getElementById('libraryformtitle').innerText = "Update library";
    document.getElementById('libraryformbutton').innerText = "Update";

    const requestOptions = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}` 
        }
    };

    fetch(url,requestOptions)
        .then(response => response.json())
        .then(data => {
            document.getElementById('admin-id').value = data.id;
            document.getElementById('admin-name').value = data.name;
            document.getElementById('admin-email').value = data.email;
            document.getElementById('password').value = data.password;
        })
        .catch(error => console.error('Error fetching data:', error));
}

function libraryAction()
{
    let url = null;
    let data = null;

    if(action=="create")
    {
        url = `http://localhost:8080/owner/library/create`;

        data = {
            name: document.getElementById('library-name').value,
            creator_id: 5,
        }
    }
    else if(action=="edit")
    {
        url = 'http://localhost:8080/owner/library/update';

        data = {
            name: document.getElementById('library-name').value,
            creator_id: 5,
            id: 1
        }
    }

    console.log(JSON.stringify(data));

    // Specify the request options
    const requestOptions = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}` 
        },
        body: JSON.stringify(data)
    };

    // Make the POST request
    fetch(url, requestOptions)
        .then(response => response.json())
        .then(data => {
            if(data.message=='success')
            {
                if(action=="create")
                {
                    alert('Library created successfully')
                    url = `http://localhost:8080/owner/library/admin/create`;

                    data = {
                        name: document.getElementById('admin-name').value,
                        email: document.getElementById('admin-email').value,
                        contact_number: 994848838,
                        role: 'Admin',
                        lib_id:2,
                        pass: document.getElementById('password').value
                    }
                }
                else if(action=="edit")
                {
                    alert('Library updated successfully')
                    url = 'http://localhost:8080/owner/library/admin/update';

                    data = {
                        name: document.getElementById('admin-name').value,
                        email: document.getElementById('admin-email').value,
                        contact_number: 994848838,
                        role: 'Admin',
                        lib_id:2,
                        pass: document.getElementById('password').value
                    }
                }

                console.log(data);

                const requestOptions = {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}` 
                    },
                    body: JSON.stringify(data)
                };

                fetch(url, requestOptions)
                    .then(response => response.json())
                    .then(data => {
                        if(data.message=='success')
                        {
                            if(action=='create')
                            {
                                alert('Admin created successfully');
                                window.location.href = 'librarylisting.html';
                            }
                            else if(action=='edit')
                            {
                                alert('Admin updated successfully');
                                window.location.href = 'librarylisting.html';
                            }  
                        }
                    })
                    .catch(error => console.error('Error fetching data:', error));
            }
        })
        .catch(error => console.error('Error fetching data:', error));
}

function Search()
{
    let search_query = document.getElementById('searchbox').value;

    const url = `http://localhost:8080/owner/library/search/${search_query}`;

    const requestOptions = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}` 
        }
    };

    fetch(url,requestOptions)
    .then(response => response.json())
    .then(data => {
        const tableBody = document.querySelector('tbody');
        tableBody.innerHTML = ''; // Clear existing table rows

        console.log(data);

        if (data && Array.isArray(data.libraries) && data.libraries.length > 0) {
            let serialNo = 1;
            data.libraries.forEach(item => {
                const row = document.createElement('tr');

                // Populate table cells
                row.innerHTML = `
                    <td>${serialNo}</td>
                    <td>${item.name}</td>
                    <td>
                        <a class="btn btn-primary" href="createlibrary.html?action=edit&id=${item.id}">Edit</a>
                        <a class="btn btn-danger" href="#" data-library-id="${item.id}" onclick="deletelibrary(event)">Delete</a>
                    </td>
                `;

                tableBody.appendChild(row);
                serialNo++;
            });
        } else {
            const row = document.createElement('tr');
            row.innerHTML = `<td colspan="8" style="text-align: center;">No librarys present</td>`;
            tableBody.appendChild(row);
        }
    })
    .catch(error => console.error('Error fetching data:', error));

}

function deletelibrary(event) 
{
    event.preventDefault();
    const libraryId = event.target.dataset.libraryId;

    if (confirm("Are you sure you want to delete this library?")) {
        // Send a request to delete the library
        fetch(`http://localhost:8080/owner/library/delete/${libraryId}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}` 
            }
        })
        .then(response => {
            if (response.ok) {
                alert("library deleted successfully");
                event.target.closest('tr').remove();
            } else {
                alert("Failed to delete the library");
            }
        })
        .catch(error => console.error('Error deleting library:', error));
    }
}

function getCookie(name) {
    const cookies = document.cookie.split('; ');
    for (let cookie of cookies) {
        const [cookieName, cookieValue] = cookie.split('=');
        if (cookieName === name) {
            return cookieValue;
        }
    }
    return null;
}

initialSetup();