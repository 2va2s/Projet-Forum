console.log("displaying thread")
let postId = window.location.href.split("/")[window.location.href.split("/").length - 1]
fetch("/posts").then(function (response) {
    return response.json()
}).then(function (response) {
    let container = document.getElementById("postContainer")
    for (let i in response) {
        if (response[i].ID == postId) {
            const div = document.createElement('div')
            fetch('../static/components/test.txt')
                .then(data => data.text())
                .then(data => {
                    data = data.split("{Pseudo}").join(response[i].Title.String).split("{Content}").join(response[i].Content).split("{Date}").join(response[i].Date).split("{PostId}").join(response[i].ID).split("{UserId}").join(response[i].UserId).split("{UpVote}").join(response[i].UpVote).split("{PostId}").join(response[i].ID)
                    // console.log("comp1 " + data)
                    fetch("/categories").then(catt => catt.json()).then(function (catt) {
                        let category = catt.filter(obj => obj.ID == response[i].Category.Int64)[0]
                        console.log(category)
                        data = data.split("{CatColor}").join(category.Color).split("{Category}").join(category.Name)
                        fetch("/users").then(users => users.json()).then(users => users.filter(user => user.ID == response[i].UserId)[0]).then(user => {
                            data = data.split("{ProfilePic}").join(user.ProfilePic)
                            div.innerHTML = data
                            div.style.width = "60vw"
                            container.appendChild(div)
                            generateChild(response, response[i], div, 60 - 10)
                        })
                        console.log("AAAA" + response)
                    })
                    // console.log(response[i].UpVote)
                })

        }
    }
}).catch(function (err) {
    console.log(err)
})

async function generateChild(response, parent, div, width) {
    console.log("div: ", response[1].ParentPostId.Int64 == parent.ID)
    let child = response.filter(post => post.ParentPostId.Int64 == parent.ID)
    console.log("child len: " + child.length)
    for (i in child) {
        childDiv = await generatePostCard(div, child[i], width)
        console.log("childDiv: " + childDiv)
        generateChild(response, child[i], childDiv, width - 10)
    }
}

async function generatePostCard(divToAdd, struc, width) {
    const div = document.createElement('div')
    let fetch1 = await fetch('../static/components/test.txt')
        .then(response => response.text())
        .then(async data => {
            data = data.split("{Pseudo}").join(struc.Title.String).split("{Content}").join(struc.Content).split("{Date}").join(struc.Date).split("{PostId}").join(struc.ID).split("{UserId}").join(struc.UserId).split("{UpVote}").join(struc.UpVote).split("{PostId}").join(struc.ID)
            // console.log("comp1 " + data)
            return await fetch("/categories").then(catt => catt.json()).then(async catt => {
                let category = catt.filter(obj => obj.ID == struc.Category.Int64)[0]
                data = data.split("{CatColor}").join(category.Color).split("{Category}").join(category.Name)
                return await fetch("/users").then(users => users.json()).then(users => users.filter(user => user.ID == struc.UserId)[0]).then(user => {
                    data = data.split("{ProfilePic}").join(user.ProfilePic)
                    div.innerHTML = data
                    div.style.width = width + "vw"
                    div.style.marginLeft = "5vw"
                    return div
                })
            })
            // console.log(response[i].UpVote)
        })
    divToAdd.appendChild(fetch1)
    return fetch1
}
