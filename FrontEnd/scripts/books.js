let bookID = null;
let action = null;

function initialSetup()
{
    const urlParams = new URLSearchParams(window.location.search);
    bookID = urlParams.get('id');
    action = urlParams.get('action');

    if(action=='create')
    {
        // Let the cotrol pass
    }
    else if(action=='edit')
    {
        getSpecificBookData();
    }
    else if(action==null)
    {
        getAllData();
    }
}

function getAllData() 
{
    const url = 'http://localhost:8080/user/books';

    fetch(url)
        .then(response => response.json())
        .then(data => {
            const tableBody = document.querySelector('tbody');
            tableBody.innerHTML = ''; // Clear existing table rows
            
            if (data.length === 0) 
            {
                const row = document.createElement('tr');
                row.innerHTML = `<td colspan="8">No books present</td>`;
                tableBody.appendChild(row);
            } 
            else 
            {
                data.forEach(item => {
                    const row = document.createElement('tr');
                    
                    // Populate table cells
                    row.innerHTML = `
                        <td>${item.isbn}</td>
                        <td>${item.title}</td>
                        <td>${item.authors}</td>
                        <td>${item.publisher}</td>
                        <td>${item.version}</td>
                        <td>${item.total_copies}</td>
                        <td>${item.available_copies}</td>
                        <td>
                            <a class="btn btn-primary" href="createbook.html?action=edit&id=${item.id}">Edit</a>
                            <a class="btn btn-danger" href="#" data-book-id="${item.id}" onclick="deleteBook(event)">Delete</a>
                        </td>
                    `;
                    
                    tableBody.appendChild(row);
                });
            }
        })
        .catch(error => console.error('Error fetching data:', error));
}

function getSpecificBookData()
{
    const url = `http://localhost:8080/admin/specific-book/${bookID}`;

    document.getElementById('bookformtitle').innerText = "Update Book";
    document.getElementById('bookformbutton').innerText = "Update";

    fetch(url)
        .then(response => response.json())
        .then(data => {
            document.getElementById('book_id').value = data.id;
            document.getElementById('lib_id').value = data.lib_id;
            document.getElementById('isbn').value = data.isbn;
            document.getElementById('title').value = data.title;
            document.getElementById('authors').value = data.authors;
            document.getElementById('publisher').value = data.publisher;
            document.getElementById('version').value = data.version;
            document.getElementById('total-copies').value = data.total_copies;
            document.getElementById('available-copies').value = data.available_copies;
        })
        .catch(error => console.error('Error fetching data:', error));
}

function bookAction(event)
{
    event.preventDefault();
    let url = null;
    let data = null;

    if(action=="create")
    {
        url = `http://localhost:8080/admin/book/create`;

        data = {
            isbn: parseInt(document.getElementById('isbn').value),
            lib_id: 11,
            title: document.getElementById('title').value,
            authors: document.getElementById('authors').value,
            publisher: document.getElementById('publisher').value,
            version: document.getElementById('version').value,
            total_copies: parseInt(document.getElementById('total-copies').value),
            available_copies: parseInt(document.getElementById('available-copies').value)
        }
    }
    else if(action=="edit")
    {
        url = 'http://localhost:8080/admin/book/update';

        data = {
            id: parseInt(document.getElementById('book_id').value),
            isbn: parseInt(document.getElementById('isbn').value),
            lib_id: parseInt(document.getElementById('lib_id').value),
            title: document.getElementById('title').value,
            authors: document.getElementById('authors').value,
            publisher: document.getElementById('publisher').value,
            version: document.getElementById('version').value,
            total_copies: parseInt(document.getElementById('total-copies').value),
            available_copies: parseInt(document.getElementById('available-copies').value)
        }
    }

    console.log(JSON.stringify(data));

    // Specify the request options
    const requestOptions = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    };

    // Make the POST request
    fetch(url, requestOptions)
        .then(response => response.json())
        .then(data => {
            if(data.message=='success')
            {
                if(action=='create')
                {
                    alert('Book created successfully');
                }
                else if(action=='edit')
                {
                    alert('Book updated successfully');
                }  
            }
        })
        .catch(error => console.error('Error fetching data:', error));
}

function Search(e)
{
    e.preventDefault();

    let search_query = document.getElementById('searchbox').value;

    const url = 'http://localhost:8080/user/book/search';

    fetch(url)
        .then(response => response.json())
        .then(data => {
            const tableBody = document.querySelector('tbody');
            tableBody.innerHTML = ''; // Clear existing table rows

            console.log(data);
            
            if (!data.books ) 
            {
                const row = document.createElement('tr');
                row.innerHTML = `<td colspan="8" style="text-align: center;">No books present</td>`;
                tableBody.appendChild(row);
            } 
            else 
            {
                data.forEach(item => {
                    const row = document.createElement('tr');
                    
                    // Populate table cells
                    row.innerHTML = `
                        <td>${item.isbn}</td>
                        <td>${item.title}</td>
                        <td>${item.authors}</td>
                        <td>${item.publisher}</td>
                        <td>${item.version}</td>
                        <td>${item.total_copies}</td>
                        <td>${item.available_copies}</td>
                        <td>
                            <a class="btn btn-primary" href="createbook.html?action=edit&id=${item.id}">Edit</a>
                            <a class="btn btn-danger" href="#">Delete</a>
                        </td>
                    `;
                    
                    tableBody.appendChild(row);
                });
            }
        })
        .catch(error => console.error('Error fetching data:', error));
}

function deleteBook(event) 
{
    event.preventDefault();
    const bookId = event.target.dataset.bookId;

    if (confirm("Are you sure you want to delete this book?")) {
        // Send a request to delete the book
        fetch(`http://localhost:8080/admin/book/delete/${bookId}`, {
            method: 'GET',
        })
        .then(response => {
            if (response.ok) {
                alert("Book deleted successfully");
                event.target.closest('tr').remove();
            } else {
                alert("Failed to delete the book");
            }
        })
        .catch(error => console.error('Error deleting book:', error));
    }
}

initialSetup();