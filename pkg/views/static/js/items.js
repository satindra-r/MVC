let counters = [];
let checkboxes = [];
let instructions = [];
let upButtons = [];
let downButtons = [];
let pageCounter = document.getElementById("page");
const pageNo = new URLSearchParams(window.location.search).get("page");
let pageup = document.getElementById("button page up");
let pagedown = document.getElementById("button page down");
let savedItems = JSON.parse(localStorage.getItem("Items")) || Object();
let savedFilters = document.getElementById("filters")["dataset"]["filters"];
let searchBox = document.getElementById("search");
const searchQuery = searchBox.value;

for (let i = 0; i < 10; i++) {
    let counter = document.getElementById("counter " + i);
    if (counter) {
        if (savedItems[counter["dataset"]["itemid"]]) {
            counter.value = savedItems[counter["dataset"]["itemid"]]["count"] || 0;
        } else {
            counter.value = 0;
        }
        counters.push(counter);
    }
    let upButton = document.getElementById("button up " + i);
    if (upButton) {
        upButtons.push(upButton);
    }
    let downButton = document.getElementById("button down " + i);
    if (downButton) {
        downButtons.push(downButton);
    }
    let checkbox = document.getElementById("checkbox " + i);
    if (checkbox) {
        if ((savedFilters & (1 << i)) !== 0) {
            checkbox.checked = true;
        }
        checkboxes.push(checkbox);
    }
    let instruction = document.getElementById("instruction " + i);
    if (instruction) {
        if (savedItems[instruction["dataset"]["itemid"]]) {
            instruction.value = savedItems[instruction["dataset"]["itemid"]]["splInstructions"] || "";
        }
        instructions.push(instruction);
    } else {
        break;
    }
}
if (counters.length === 0) {
    if (parseInt(pageCounter.value) !== 1) {
        document.location.replace("/items?page=" + Math.max(0, parseInt(pageCounter.value - 1 || "1")) + "&filters=" + savedFilters + "&search=" + searchQuery);
    }
}
for (let i = 0; i < counters.length; i++) {
    counters[i].addEventListener("input", function (e) {
        let counterValue = Math.max(0, parseInt(e.target.value || "0"));
        e.target.value = counterValue;
        let items = JSON.parse(localStorage.getItem("Items")) || Object();
        if (!items[e.target["dataset"]["itemid"]]) {
            items[e.target["dataset"]["itemid"]] = Object();
        }
        if (counterValue) {
            items[e.target["dataset"]["itemid"]]["count"] = counterValue;
        } else {
            delete items[e.target["dataset"]["itemid"]];
        }
        localStorage.setItem("Items", JSON.stringify(items));
    })
}

for (let i = 0; i < upButtons.length; i++) {
    upButtons[i].addEventListener("click", function (e) {
        let counterValue = Math.max(0, parseInt(counters[i].value || "0") + 1);
        counters[i].value = counterValue;
        let items = JSON.parse(localStorage.getItem("Items")) || Object();
        if (!items[e.target["dataset"]["itemid"]]) {
            items[e.target["dataset"]["itemid"]] = Object();
        }
        if (counterValue) {
            items[e.target["dataset"]["itemid"]]["count"] = counterValue;
        } else {
            delete items[e.target["dataset"]["itemid"]];
        }
        localStorage.setItem("Items", JSON.stringify(items));
    })
}

for (let i = 0; i < downButtons.length; i++) {
    downButtons[i].addEventListener("click", function (e) {
        let counterValue = Math.max(0, parseInt(counters[i].value || "0") - 1);
        counters[i].value = counterValue;
        let items = JSON.parse(localStorage.getItem("Items")) || Object();
        if (!items[e.target["dataset"]["itemid"]]) {
            items[e.target["dataset"]["itemid"]] = Object();
        }
        if (counterValue) {
            items[e.target["dataset"]["itemid"]]["count"] = counterValue;
        } else {
            delete items[e.target["dataset"]["itemid"]];
        }
        localStorage.setItem("Items", JSON.stringify(items));
    })
}

for (let i = 0; i < instructions.length; i++) {
    instructions[i].addEventListener("input", function (e) {
        let instructionValue = e.target.value;
        let items = JSON.parse(localStorage.getItem("Items")) || Object();
        if (!items[e.target["dataset"]["itemid"]]) {
            items[e.target["dataset"]["itemid"]] = Object();
        }
        if (instructionValue) {
            items[e.target["dataset"]["itemid"]]["splInstructions"] = instructionValue;
            if ((items[e.target["dataset"]["itemid"]]["count"] || 0) === 0) {
                items[e.target["dataset"]["itemid"]]["count"] = 1;
                counters[i].value = 1
            }
        } else {
            items[e.target["dataset"]["itemid"]]["splInstructions"] = "";
        }
        localStorage.setItem("Items", JSON.stringify(items));
    })
}

for (let i = 0; i < checkboxes.length; i++) {
    checkboxes[i].addEventListener("click", function (e) {
        savedFilters ^= (1 << (e.target["dataset"]["sectionid"] - 1));
        document.location.href = "/items?page=" + pageCounter.value + "&filters=" + savedFilters + "&search=" + searchQuery;
    })
}

pageup.addEventListener("click", function () {
    document.location.href = "/items?page=" + (parseInt(pageCounter.value) + 1) + "&filters=" + savedFilters + "&search=" + searchQuery;
})

pagedown.addEventListener("click", function () {
    if (parseInt(pageCounter.value) > 1) {
        document.location.href = "/items?page=" + (parseInt(pageCounter.value) - 1) + "&filters=" + savedFilters + "&search=" + searchQuery;
    }
})

pageCounter.addEventListener("keyup", function (e) {
    if (e.key === "Enter") {
        if (e.target.value !== pageNo) {
            document.location.href = "/items?page=" + e.target.value + "&filters=" + savedFilters + "&search=" + searchQuery;
        }
    } else {
        e.target.value = Math.max(1, parseInt(e.target.value || "1"));
    }
})

searchBox.addEventListener("keydown", function (e) {
    if (e.key === "Enter") {
        if (searchBox.value !== searchQuery) {
            document.location.href = "/items?page=" + pageCounter.value + "&filters=" + savedFilters + "&search=" + e.target.value;
        }
    }
})

document.getElementById("order").addEventListener("click", async function () {
    savedItems = JSON.parse(localStorage.getItem("Items"));
    let orderedItems = []
    for (let i in savedItems) {
        if (savedItems[i]["count"] > 0) {
            orderedItems.push({
                "itemId": parseInt(i),
                "splInstructions": savedItems[i]["splInstructions"],
                "count": savedItems[i]["count"]
            })
        }
    }
    let res = await fetch("/api/order", {
        method: "POST", headers: {"Content-Type": "application/json",}, body: JSON.stringify({"Items": orderedItems}),
    })
    if (res["ok"]) {
        localStorage.setItem("Items", "{}");
        document.location.href = "/orders";
    }
})

document.getElementById("logout").addEventListener("click", function () {
    localStorage.clear();
    document.cookie = 'JWT=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
});