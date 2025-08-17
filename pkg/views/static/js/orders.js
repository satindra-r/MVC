let dropdowns = [];
let pays = [];
let upButtons = [];
let downButtons = [];
let pageCounter = document.getElementById("page");
const pageNo = new URLSearchParams(window.location.search).get("page");
let pageup = document.getElementById("button page up");
let pagedown = document.getElementById("button page down");
for (let i = 0; i < 10; i++) {
    let pay = document.getElementById("pay " + i);
    if (pay) {
        pays.push(pay);
    }
    for (let j = 0; ; j++) {
        let upButton = document.getElementById("button up " + i + " " + j);
        if (upButton) {
            upButtons.push(upButton);
        }
        let downButton = document.getElementById("button down " + i + " " + j);
        if (downButton) {
            downButtons.push(downButton);
        } else {
            break;
        }
    }
    let dropdown = document.getElementById("dropdown " + i);
    if (dropdown) {
        dropdowns.push(dropdown);
    } else {
        break;
    }
}

if (dropdowns.length === 0 && pageNo > 1) {
    document.location.replace("/orders?page=" + Math.max(0, parseInt(pageCounter.value - 1 || "1")));
}

for (let i = 0; i < dropdowns.length; i++) {
    dropdowns[i].addEventListener("click", function (e) {
        let toggleElements = document.getElementsByClassName("hidden-" + i);
        for (let j = 0; j < toggleElements.length; j++) {
            toggleElements[j].classList.toggle("hidden");
        }
        if (e.target.innerText === "v") {
            e.target.innerText = ">";
        } else {
            e.target.innerText = "v";
        }
    });
}

for (let i = 0; i < pays.length; i++) {
    pays[i].addEventListener("click", function (e) {
        document.location.href = "/bill?order=" + e.target["dataset"]["orderid"];
    });
}

for (let i = 0; i < upButtons.length; i++) {
    upButtons[i].addEventListener("click", async function (e) {
        let res = await fetch("/api/dish/count", {
            method: "PUT",
            headers: {"Content-Type": "application/json", "dishId": e.target["dataset"]["dishid"], "count": 1,},
        })
        if (res["ok"]) {
            localStorage.setItem("Items", "{}");
            document.location.href = "/orders?page=" + pageCounter.value;
        }
    })
}


for (let i = 0; i < downButtons.length; i++) {
    downButtons[i].addEventListener("click", async function (e) {
        let res = await fetch("/api/dish/count", {
            method: "PUT",
            headers: {"Content-Type": "application/json", "dishId": e.target["dataset"]["dishid"], "count": -1,},
        })
        if (res["ok"]) {
            localStorage.setItem("Items", "{}");
            document.location.href = "/orders?page=" + pageCounter.value;
        }
    })
}

pageup.addEventListener("click", function () {
    document.location.href = "/orders?page=" + (parseInt(pageCounter.value) + 1);
})

pagedown.addEventListener("click", function () {
    if (parseInt(pageCounter.value) > 1) {
        document.location.href = "/orders?page=" + (parseInt(pageCounter.value) - 1);
    }
})

pageCounter.addEventListener("keyup", function (e) {
    if (e.key === "Enter") {
        if (e.target.value !== pageNo) {
            document.location.href = "/orders?page=" + e.target.value;
        }
    } else {
        e.target.value = Math.max(1, parseInt(e.target.value || "1"));
    }
})

document.getElementById("logout").addEventListener("click", function (e) {
    localStorage.clear();
    document.cookie = 'JWT=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
});