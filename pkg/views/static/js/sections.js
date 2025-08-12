let upButtons = [];
let downButtons = [];
let delButtons = [];
let addButton = document.getElementById("button add")
let error = document.getElementById("Error");

for (let i = 0; i < 10; i++) {
    let upButton = document.getElementById("button up " + i);
    if (upButton) {
        upButtons.push(upButton);
    }
    let downButton = document.getElementById("button down " + i);
    if (downButton) {
        downButtons.push(downButton);
    }
    let delButton = document.getElementById("button del " + i);
    if (delButton) {
        delButtons.push(delButton);
    } else {
        break;
    }
}

upButtons[0].setAttribute("disabled", "");
downButtons[downButtons.length - 1].setAttribute("disabled", "");

addButton.addEventListener("click", async function () {
    let res = await fetch("/api/sections", {
        method: "POST",
        headers: {
            "SectionName": document.getElementById("SectionName").value,
        }
    })
    if (res["ok"]) {
        location.reload();
    } else {
        error.classList.remove("hidden");
        error.innerText = await res.text()
    }
})

for (let i = 0; i < upButtons.length; i++) {
    upButtons[i].addEventListener("click", async function (e) {
        let res = await fetch("/api/sections", {
            method: "PUT",
            headers: {
                "SectionId1": e.target["dataset"]["sectionid"],
                "SectionId2": upButtons[i - 1]["dataset"]["sectionid"],
            }
        })
        if (res["ok"]) {
            location.reload();
        } else {
            error.classList.remove("hidden");
            error.innerText = await res.text()
        }
    })
}
for (let i = 0; i < downButtons.length; i++) {
    downButtons[i].addEventListener("click", async function (e) {
        let res = await fetch("/api/sections", {
            method: "PUT",
            headers: {
                "SectionId1": e.target["dataset"]["sectionid"],
                "SectionId2": downButtons[i + 1]["dataset"]["sectionid"],
            }
        })
        if (res["ok"]) {
            location.reload();
        } else {
            error.classList.remove("hidden");
            error.innerText = await res.text()
        }
    })
}

for (let i = 0; i < delButtons.length; i++) {
    delButtons[i].addEventListener("click", async function (e) {
        let res = await fetch("/api/sections", {
            method: "DELETE",
            headers: {
                "SectionId": e.target["dataset"]["sectionid"],
            }
        })
        if (res["ok"]) {
            location.reload();
        } else {
            error.classList.remove("hidden");
            error.innerText = await res.text()
        }
    })
}

document.getElementById("logout").addEventListener("click", function () {
    localStorage.clear();
    document.cookie = 'JWT=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
});