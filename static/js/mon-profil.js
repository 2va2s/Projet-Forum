let monCompte = document.getElementById("monCompte");
let posts = document.getElementById("posts");
let paramètres = document.getElementById("paramètes");
let déconnexion = document.getElementById("déconnexion");
let photo = document.getElementById("photo");
let pseudo = document.getElementById("pseudo");
let mail = document.getElementById("e-mail");
let rang = document.getElementById("rang");
let liste = document.getElementById("liste");
let listePosts = document.getElementsByClassName("listePosts");
let déco = document.getElementById("déco");
let mesCommentaires = document.getElementById("mesCommentaires");
let commentaires = document.getElementById("commentaires");
let badges = document.getElementById("badges");
let déroul = document.getElementById("scroll");
let non = document.getElementById("non")
let infoDiv = document.getElementById("info")

monCompte.addEventListener("click", () => {


        console.log("hvjyg")
        infoDiv.style.backgroundColor = "#1C1F4C"
        photo.style.display = "block";
        pseudo.style.display = "block";
        mail.style.display = "block";
        rang.style.display = "block"
        badges.style.display = "block"
        liste.style.display = "none"
        déco.style.display = "none";
        commentaires.style.display = "none"
        déroul.style.display = "none"
        openPopup.style.display = "none"

}
)

posts.addEventListener("click", () => {


        déroul.style.display = "block"
        photo.style.display = "none";
        pseudo.style.display = "none";
        mail.style.display = "none";
        rang.style.display = "none";
        déco.style.display = "none";
        commentaires.style.display = "none"
        badges.style.display = "none"
        openPopup.style.display = "none"
        infoDiv.style.backgroundColor = "#698EA2"


}
)

déconnexion.addEventListener("click", () => {
        openPopup.style.display = "block"

})

mesCommentaires.addEventListener("click", () => {
        déroul.style.display = "block"
        déco.style.display = "none"
        liste.style.display = "none"
        photo.style.display = "none";
        pseudo.style.display = "none";
        mail.style.display = "none";
        rang.style.display = "none";
        badges.style.display = "none"
        openPopup.style.display = "none"
        infoDiv.style.backgroundColor = "#698EA2"


})

non.addEventListener("click", () => {
        openPopup.style.display = "none"
})

fetch("/cookies-data").then((response) => response.json()).then(data => {
        console.log("datas: " + data.pseudo)
        document.getElementById("PseudoDisplay").innerText = data.pseudo
        fetch("/users").then(response => response.json()).then(response => {
                let userData = response.filter(user => user.ID == data.user_id)[0]
                document.getElementById("EmailDisplay").innerText = userData.Mail
                document.getElementById("NumberDisplay").innerText = userData.Number
                document.getElementById("photo").src = "../static/img/avatar/" + userData.ProfilePic + ".png"
        })
})

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

        const upvotes = await fetch("/upvote").then(data => data.json()).then(async data => {
                let test = data.filter(upvote => {
                        console.log(upvote.UserId, response.ID)
                        return upvote.UserId == response.ID
                })
                console.log(test)
                test = test.map(value => value.PostId)
                console.log(test)

                // console.log(upvotes.UserID)
                return await fetch("/posts").then(dat => dat.json()).then(poss => poss.filter(pos => {
                        return test.includes(String(pos.ID))
                }))
        })


        console.log(upvotes)
        upvotes.forEach(upvote => {
                console.log("okokok")
                let upvoteDiv = document.getElementById("postUpvotes")
                let newDiv = document.createElement("div")
                newDiv.innerText = upvote.Content
                // newDiv2.innerText = upvote.Title.String

                // newDiv.innerText = upvote.UserID.String
                newDiv.style.cssText = "margin-top: 5%; cursor:pointer;border-top: 0.1rem solid; background-color:#ffffff;height:20%"
                newDiv.addEventListener("click", function () { location.href = "/topic/" + upvote.ID })

                upvoteDiv.appendChild(newDiv)
        })
})