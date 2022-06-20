fetch('/users').then(function (response) {
    return response.json()
}).then(function (response) {
    console.log(window.location.href)
    const url = window.location.href.split('/')
    return response.filter(user => user.ID == url[url.length - 1])[0]
}).then(async response => {
    console.log("response: " + response.Pseudo)
    document.getElementById("profilname").innerText = response.Pseudo
    document.getElementById("profilrank").innerText = (response.Level == "1") ? "Utilisateur" : (response.Level == "2") ? "Modérateur" : "Administrateur"
    let pp = document.getElementById("profilavatar")
    pp.src = "../static/img/avatar/" + response.ProfilePic + ".png"
    const posts = await fetch("/posts").then(data => data.json()).then(data => data.filter(post => post.UserId == response.ID))
    console.log("ses posts: " + posts)
    posts.forEach(post => {
        let postsDiv = document.getElementById("profilpostedposts")
        let newDiv = document.createElement("div")
        newDiv.innerText = post.Title.String
        newDiv.style.cssText = "margin-top: 5%; cursor:pointer;border-top: 0.1rem solid; background-color:#ffffff;height:20%"
        newDiv.addEventListener("click", function () { location.href = "/topic/" + post.ID })
        postsDiv.appendChild(newDiv)
    })
    // const upvotes =
    // const reponses = 
})