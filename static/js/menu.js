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
  })
}