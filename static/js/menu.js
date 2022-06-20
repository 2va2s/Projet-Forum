const menucontent = document.querySelector('#connectedmenucontent')


function hideMenu() {
  menucontent.classList.remove('open')
}

function showContent() {
  menucontent.classList.add('open')
}

function onClickRegister() {
  fetch("/signin", {
    method: 'POST',
    headers: {
      "content-type": "application/json"
    },
    body: JSON.stringify({
      Pseudo: document.getElementById("signinPseudo").value,
      Email: document.getElementById("signinEmail").value,
      Number: document.getElementById("signinNumber").value,
      Password: document.getElementById("signinPassword").value,
      Password2: document.getElementById("signinPassword2").value
    })
  }).then(response => response.json())
}

function onClickLogin() {
  fetch("/login", {
    method: 'POST',
    headers: {
      "content-type": "application/json"
    },
    body: JSON.stringify({
      Pseudo: document.getElementById("loginPseudo").value,
      Password: document.getElementById("loginPassword").value
    })
  }).then(response => response.text())
}

async function getUserId() {
  connectedUser = await fetch('/cookies-data').then(response => response.json()).then(response => response.user_id).catch(response => response)
  console.log("aa" + connectedUser)
  return connectedUser
}
console.log("getUserId: " + getUserId())

async function loadVote(postId) {
  const connectedUser = await getUserId()
  console.log("loading upvote")
  fetch("/IsUpvoted", {
    method: "POST",
    headers: {
      "content-type": "application/json"
    },
    body: JSON.stringify({
      UserId: connectedUser,
      PostId: postId.toString()
    })
  }).then(response => response.text()).then(response => document.getElementById("arrowUp#" + postId).className = response)
}

async function vote(id) {
  const connectedUser = await getUserId()
  fetch("/UpdateVote", {
    method: "POST",
    headers: {
      "content-type": "application/json"
    },
    body: JSON.stringify({
      Table: "post",
      Id: connectedUser,
      Field: "UpVote",
      Where: "ID",
      PostId: id.toString(),
    })
  }).then(response => response.text()).then(data => {
    if (data != "isntLogged") {
      arrowUp = document.getElementById("arrowUp#" + id)
      if (data > document.getElementById("upvoteNumber#" + id).innerText) {
        arrowUp.className = "upvoteUp"
      } else {
        arrowUp.className = "upvote"
      }
      // console.log("data:" + data)
      document.getElementById("upvoteNumber#" + id).innerText = data
    } else {
      console.log("notLogged")
      document.getElementById("notLoggedAlert").style.display = "flex"
    }
  })
}

fetch("/cookies-data").then(response => response.json()).then(data => {
  if (data != null) console.log("bushi lz")
  if (data != "") {
    document.getElementById("connected-menu").style.display = 'block'
    document.getElementById("signin/login").style.display = 'none'
    document.getElementById("DisplayPseudo").innerText = data.pseudo
  }
}).catch(response => response)

const redirectToPost = (id) => {
  fetch('/posts').then(location.href = '/topic/' + id)
}

const redirectToUser = (id) => {
  fetch('/users').then(location.href = '/user/' + id)
}
