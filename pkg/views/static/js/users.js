let userRadios = [];
let chefRadios = [];
let adminRadios = [];

let pageCounter = document.getElementById("page");
const pageNo = new URLSearchParams(window.location.search).get("page");
let pageup = document.getElementById("button page up");
let pagedown = document.getElementById("button page down");

for (let i = 0; i < 10; i++) {

    let userRadio = document.getElementById("radio 1 " + i);
    if (userRadio) {
        userRadios.push(userRadio);
    }

    let chefRadio = document.getElementById("radio 2 " + i);
    if (chefRadio) {
        chefRadios.push(chefRadio);
    }

    let adminRadio = document.getElementById("radio 3 " + i);
    if (adminRadio) {
        adminRadios.push(adminRadio);
    } else {
        break;
    }
}


if (userRadios.length === 0 && pageNo > 1) {
    document.location.replace("/users?page=" + Math.max(0, parseInt(pageCounter.value - 1 || "1")));
}

for (let i = 0; i < userRadios.length; i++) {
    userRadios[i].addEventListener("click", async function (e) {
        let res = await fetch("/api/user", {
            method: "PUT", headers: {
                UserId: e.target["dataset"]["userid"], Role: "User"
            }
        });
        if (res["ok"]) {
            location.reload();
        }
    })

}

for (let i = 0; i < chefRadios.length; i++) {
    chefRadios[i].addEventListener("click", async function (e) {
        let res = await fetch("/api/user", {
            method: "PUT", headers: {
                UserId: e.target["dataset"]["userid"], Role: "Chef"
            }
        });
        if (res["ok"]) {
            location.reload();
        }
    })

}

for (let i = 0; i < adminRadios.length; i++) {
    adminRadios[i].addEventListener("click", async function (e) {
        let res = await fetch("/api/user", {
            method: "PUT", headers: {
                UserId: e.target["dataset"]["userid"], Role: "Admin"
            }
        });
        if (res["ok"]) {
            location.reload();
        }
    })

}

pageup.addEventListener("click", function () {
    document.location.href = "/users?page=" + (parseInt(pageCounter.value) + 1);
})

pagedown.addEventListener("click", function () {
    if (parseInt(pageCounter.value) > 1) {
        document.location.href = "/users?page=" + (parseInt(pageCounter.value) - 1);
    }
})

pageCounter.addEventListener("keyup", function (e) {
    if (e.key === "Enter") {
        if (e.target.value !== pageNo) {
            document.location.href = "/users?page=" + e.target.value;
        }
    } else {
        e.target.value = Math.max(1, parseInt(e.target.value || "1"));
    }
})

document.getElementById("logout").addEventListener("click", function () {
    localStorage.clear();
    document.cookie = 'JWT=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
});