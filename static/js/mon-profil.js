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
let openPopup = document.getElementById("openPopup")

monCompte.addEventListener("click", () => {
    
    
        console.log("hvjyg")
        photo.style.display = "block";
        pseudo.style.display = "block";
        mail.style.display = "block";
        rang.style.display = "block"
        badges.style.display = "block"
        liste.style.display = "none"
        déco.style.display = "none";
        commentaires.style.display = "none"
        déroul.style.display = "none"
        openPopup.style.display ="none"
          
    }
)

posts.addEventListener("click", () => {
    
    
        déroul.style.display = "block"
        photo.style.display = "none";
        pseudo.style.display = "none";
        mail.style.display = "none";
        rang.style.display = "none";
        déco.style.display = "none";
        commentaires.style.display ="none"
        badges.style.display ="none"
        openPopup.style.display ="none"
        
    } 
)

déconnexion.addEventListener("click", () => {
        openPopup.style.display ="block"
        
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
        openPopup.style.display ="none"
        
})
