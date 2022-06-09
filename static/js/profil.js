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
let déco = document.getElementById("déco")

monCompte.addEventListener("click", () => {
    
    
        console.log("hvjyg")
        photo.style.display = "block";
        pseudo.style.display = "block";
        mail.style.display = "block";
        rang.style.display = "block"
        liste.style.display = "none"
        déco.style.display = "none"
   
         
        
        
    }
)

posts.addEventListener("click", () => {
    
    
        liste.style.display = "block"
        photo.style.display = "none";
        pseudo.style.display = "none";
        mail.style.display = "none";
        rang.style.display = "none";
        déco.style.display = "none"
    } 
)

déconnexion.addEventListener("click", () => {
        déco.style.display = "block"
        liste.style.display = "none"
        photo.style.display = "none";
        pseudo.style.display = "none";
        mail.style.display = "none";
        rang.style.display = "none";
        
})

