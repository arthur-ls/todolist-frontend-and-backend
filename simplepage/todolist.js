var myNodeList = document.getElementsByTagName("LI");
var i;
for (i = 0; i < myNodeList.length; i++) {
    var spanClose = document.createElement("SPAN");
    var spanEdit = document.createElement("SPAN");
    var spanCheck = document.createElement("SPAN");
    var spanUndo = document.createElement("SPAN");
    var txtClose = document.createTextNode("\u00D7");
    var txtEdit = document.createTextNode("\u2710");
    var txtCheck = document.createTextNode("\u2713");
    var txtUndo = document.createTextNode("\u21BB");
    spanClose.className = "close";
    spanEdit.className = "edit";
    spanCheck.className = "check";
    spanUndo.className = "undo";
    spanClose.appendChild(txtClose);
    spanEdit.appendChild(txtEdit);
    spanCheck.appendChild(txtCheck);
    spanUndo.appendChild(txtUndo);
    myNodeList[i].appendChild(spanClose);
    myNodeList[i].appendChild(spanEdit);
    myNodeList[i].appendChild(spanCheck);
    myNodeList[i].appendChild(spanUndo);
}

var close = document.getElementsByClassName("close");
var i;
for (i = 0; i < close.length; i++) {
    close[i].onclick = function() {
        var div = this.parentElement;
        div.style.display = "none"
    }
}

var check = document.getElementsByClassName("check");
var i;
for (i = 0; i < check.length; i++) {
    check[i].onclick = function() {
        var div = this.parentElement;
        div.style.display = "none"
    }
}

// var list = document.getElementsByClassName("check");
// list.addEventListener('click', function(event) {
//     if (event.target.className === "check") {
//         //event.target.classList.toggle('checked');
//     }
// }, false)

var undo = document.getElementsByClassName("undo");

var edit = document.getElementsByClassName("edit");

function newElement() {
    var li = document.createElement("li");
    var inputValue = document.getElementById("todolistInput").value;
    var t = document.createTextNode(inputValue);
    li.appendChild(t);
    if (inputValue === '') {
        alert("You must right something!");
    } else {
        document.getElementById("todoList").appendChild(li);
    }
    document.getElementById("todolistInput").value = "";

    var spanClose = document.createElement("SPAN");
    var spanCheck = document.createElement("SPAN");
    var spanEdit = document.createElement("SPAN");
    var spanUndo = document.createElement("SPAN");
    var txtClose = document.createTextNode("\u2717");
    var txtCheck = document.createTextNode("\u2713");
    var txtEdit = document.createTextNode("\u2710");
    var txtUndo = document.createTextNode("\u21BB");
    spanClose.className = "close";
    spanCheck.className = "check";
    spanEdit.className = "edit";
    spanUndo.className = "undo";
    spanClose.appendChild(txtClose);
    spanEdit.appendChild(txtEdit);
    spanCheck.appendChild(txtCheck);
    spanUndo.appendChild(txtUndo);
    li.appendChild(spanClose);
    li.appendChild(spanEdit);
    li.appendChild(spanCheck);
    li.appendChild(spanUndo);

    for (i = 0; i < close.length; i++) {
        close[i].onclick = function() {
            var div = this.parentElement;
            div.style.display = "none";
            localStorage.setItem('delete', div.innerText.replace(/[^a-zA-Z ]/g, ""));
            let data = {
                id: null,
                Description: document.getElementsByClassName('close')[0].parentElement.innerText.replace(/[^a-zA-Z ]/g, ""),
                Status: document.getElementsByClassName('close')[0].className
            };
            fetch("http://localhost:3000/todo", {
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                method: "POST",
                body: JSON.stringify(data)
            }).then(response => {
                console.log("Request complete! response:", response.body);
            });
        }
    }

    for (i = 0; i < check.length; i++) {
        check[i].onclick = function() {
            var div = this.parentElement;
            div.style.backgroundColor = "#52c64c";
            div.style.textDecoration = "line-through";
            localStorage.setItem('checked', div.innerText.replace(/[^a-zA-Z ]/g, ""));
            let data = {
                id: null,
                Description: document.getElementsByClassName('check')[0].parentElement.innerText.replace(/[^a-zA-Z ]/g, ""),
                Status: document.getElementsByClassName('check')[0].className
            };
            fetch("http://localhost:3000/todo", {
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                method: "POST",
                body: JSON.stringify(data)
            }).then(response => {
                console.log("Request complete! response:", response.body);
            });
        }
    }

    for (i = 0; i < undo.length; i++) {
        undo[i].onclick = function() {
            var div = this.parentElement;
            div.style.backgroundColor = "#f9f9f9";
            div.style.textDecoration = "none";
            localStorage.setItem('undo', div.innerText.replace(/[^a-zA-Z ]/g, ""));
            let data = {
                id: null,
                Description: document.getElementsByClassName('undo')[0].parentElement.innerText.replace(/[^a-zA-Z ]/g, ""),
                Status: document.getElementsByClassName('undo')[0].className
            };
            fetch("http://localhost:3000/todo", {
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                method: "POST",
                body: JSON.stringify(data)
            }).then(response => {
                console.log("Request complete! response:", response.body);
            });
        }
    }

    for (i = 0; i < edit.length; i++) {
        edit[i].onclick = function() {
            var div = this.parentElement;
            var p = prompt("Edit Task");
            div.innerHTML = p;
            localStorage.setItem('edit', div.innerText.replace(/[^a-zA-Z ]/g, ""));
            li.appendChild(spanClose);
            li.appendChild(spanEdit);
            li.appendChild(spanCheck);
            li.appendChild(spanUndo);
            let data = {
                id: null,
                Description: document.getElementsByClassName('edit')[0].parentElement.innerText.replace(/[^a-zA-Z ]/g, ""),
                Status: document.getElementsByClassName('edit')[0].className
            };
            fetch("http://localhost:3000/todo", {
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                method: "POST",
                body: JSON.stringify(data)
            }).then(response => {
                console.log("Request complete! response:", response.body);
            });
        }
    }
}