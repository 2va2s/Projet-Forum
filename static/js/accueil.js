let conn = document.getElementById("connected")
if (conn.innerHTML == " Connecté en tant que [] ") {
    conn.innerHTML = "Non connecté"
    document.getElementById("logout").style.display = 'none';
    document.getElementById("login").style.display = 'block';
}



export const displayPosts = () => {
    fetch("/topics").then(function (response) {
        return response.json()
    }).then(function (response) {
        let list = document.getElementById("topicList")
        for (let i in response) {
            const p = document.createElement('div')
            fetch('./static/components/test.txt')
                .then(response => response.text())
                .then(data => {
                    // Do something with your data
                    data = data.split("{Pseudo}").join(response[i].Title.String).split("{Content}").join(response[i].Content).split("{Date}").join(response[i].Date).split("{PostId}").join(response[i].ID).split("{UserId}").join(response[i].UserId).split("{UpVote}").join(response[i].UpVote)
                    // console.log("comp1 " + data)
                    fetch("/categories").then(catt => catt.json()).then(function (catt) {
                        let category = catt.filter(obj => obj.ID == response[i].Category.Int64)[0]
                        console.log(category)
                        data = data.split("{CatColor}").join(category.Color).split("{Category}").join(category.Name)
                        p.innerHTML = data
                        list.appendChild(p)
                    })
                    // console.log(response[i].UpVote)
                })
        }
    }).catch(function (err) {
        console.log(err)
    })
}
