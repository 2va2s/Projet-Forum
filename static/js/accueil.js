fetch("/posts").then(function (response) {
    return response.json()
}).then(function (response) {
    let list = document.getElementById("topicList")
    for (let i in response) {
        list.appendChild(document.createTextNode(response[i].Content) + "[" + response[i].Date + "]")
        list.appendChild(document.createElement("br"))
    }
}).catch(function (err) {
    console.log(err)
})

let conn = document.getElementById("connected")
if (conn.innerHTML == " Connecté en tant que [] ") {
    conn.innerHTML = "Non connecté"
    document.getElementById("logout").style.display = 'none';
    document.getElementById("login").style.display = 'block';
}



export const newPost = (post) => {
    const p = document.createElement('div')
    p.innerHTML = 
        `
            <div class="toto">
                <p>${post.title}</p>
            </div>
        `
    return p
}










