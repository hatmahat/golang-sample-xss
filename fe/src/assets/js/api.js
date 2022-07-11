async function getUsers() {
    let url =  'http://localhost:8888/api/user';
    try {
        let res = await fetch(url);
        console.log(res.data)
        return await res.json()
    } catch (error) {
        console.log(error)
    }
}

async function renderUsers() {
    let users = await getUsers();
    let html = '';
    users.forEach(user => {
        console.log(user)
        let htmlSegment = `<div>
    ${user.firstName} ${user.lastName}
    </div>`;
        html += htmlSegment
    });

    let container = document.querySelector('.container');
    container.innerHTML = html;
}

renderUsers();