var upvotenumber = document.getElementById("upvotenumber")
var upvotecounter = 0

upvotenumber.innerHTML = upvotecounter

function upvoteFunc(){
  upvotecounter += 1
  upvotenumber.innerHTML = ""
  upvotenumber.innerHTML = upvotecounter
}

function downFunc(){
    upvotecounter -= 1
    upvotenumber.innerHTML = ""
    upvotenumber.innerHTML = upvotecounter
  }

var upvote = document.getElementById('upvote');
upvote.addEventListener('click', upvoteFunc);

var downvote = document.getElementById('downvote');
downvote.addEventListener('click', downFunc);