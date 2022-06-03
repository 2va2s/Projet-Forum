let conn = document.getElementById("connected")
if (conn.innerHTML == " Connecté en tant que [] ") {
    conn.innerHTML = "Non connecté"
    document.getElementById("logout").style.display = 'none';
    document.getElementById("login").style.display = 'block';
}



export const displayPosts = () => {
    fetch("/posts").then(function (response) {
        return response.json()
    }).then(function (response) {
        let list = document.getElementById("topicList")
        for (let i in response) {
            const p = document.createElement('div')
            fetch('./static/components/test.txt')
                .then(response => response.text())
                .then(data => {
                    // Do something with your data
                    data = data.split("{Pseudo}").join(response[i].Title.String).split("{Content}").join(response[i].Content).split("{Date}").join(response[i].Date).split("{PostId}").join(response[i].ID).split("{UserId}").join(response[i].UserId)
                    console.log("comp1 " + data)
                    p.innerHTML = data
                    list.appendChild(p)
                })
        }
    }).catch(function (err) {
        console.log(err)
    })
}
