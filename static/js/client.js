const timerDelay = 1000;

let timerList;

var socket, userId = -1;

window.onload = function () {
    loadAllData();
    loadButtonHandlers();
    connectSocket();
};

function loadButtonHandlers() {
    let buttons = document.querySelectorAll("input[id^='id-button']");
    timerList = new Array(buttons.length);
    for (let button of buttons) {
        let id = button.id.split('id-button-id-')[1];
        let parent = button.parentNode;
        button.onclick = function () {
            if (parent.classList.contains("timer-run")) {
                socket.send(constructSocketData("StopTimer", id));
                stopTimer(id);
            } else {
                socket.send(constructSocketData("StartTimer", id));
                startTimer(id);
            }
        }
    }
}

function updateTimer(timerText) {
    let timeArray = timerText.innerText.split(':');
    timeArray[2] = parseInt(timeArray[2]) + 1;
    if (timeArray[2] >= 60) {
        timeArray[2] = 0;
        timeArray[1] = parseInt(timeArray[1]) + 1;
        if (timeArray[1] >= 60) {
            timeArray[2] = 0;
            timeArray[0] = parseInt(timeArray[0]) + 1;
        }
    }
    let date = new Date(0, 1, 1, timeArray[0], timeArray[1], timeArray[2]);
    timerText.textContent = date.toTimeString().split(' GMT')[0];
}

function constructSocketData(action, value) {
    return JSON.stringify({
        "UserId": userId,
        "Action": action,
        "Value" : value
    })
}

function startTimer(id) {
    let button = document.getElementById("id-button-id-" + id);
    let parent = button.parentNode;
    parent.classList.add("timer-run");
    button.value = "Остановить";
    timerList[id] = setInterval(function () {
        let elem = document.querySelector("span[id='id-timer-id-" + id + "'");
        updateTimer(elem);
    }, timerDelay);
}

function stopTimer(id) {
    let button = document.getElementById("id-button-id-" + id);
    let parent = button.parentNode;
    parent.classList.remove("timer-run");
    button.value = "Запустить";
    clearInterval(timerList[id]);
}

function loadAllData() {
    let timers = document.querySelectorAll("span[id^='id-timer']");
    for (let timer of timers) {
        id = timer.id.split("id-timer-id-")[1];
        time = loadTodayTime(id).split(' +')[0];
        console.log("Timer-" + id + " loaded time: " + time);
        timer.textContent = new Date(time).toTimeString().split(' GMT')[0];
    }
}

function loadTodayTime(idTimer) {
    let xhr = new XMLHttpRequest();
    xhr.open('GET', '/api/getTimeToday?idTimer=' + idTimer, false);
    xhr.send();
    if (xhr.status === 200) {
        console.log("Good loadTimeToday()");
        return xhr.responseText;
    } else {
        console.log("loadTimeToday() xhr status is not 200");
        alert(xhr.status + ': ' + xhr.statusText);
        return null;
    }
}

function connectSocket() {
    socket = new WebSocket("ws:localhost:8593/connect");

    socket.onopen = function () {
        console.log("Web socket connected.");
    };

    socket.onclose = function (event) {
        if (event.wasClean) {
            console.log("WebSocket connection closed clean")
        } else {
            console.log("WebSocket connection closed not clean")
        }
        console.log("WebSocket code:" + event.code + " with reason: " + event.reason)
    };

    socket.onmessage = function (event) {
        console.log("WebSocket got message:" + event.data);
        json = JSON.parse(event.data);
        // Todo: Sometimes from server come package without UserID
        /*if(json.UserId !== userId) {
            console.log("I got not my message")
        } else {
            handleAction(json.Action, json.Value)
        }*/
        handleAction(json.Action, json.Value)
    };

    socket.onerror = function (error) {
        console.log("WebSocket error:" + error)
    };
}

function handleAction(Action, Value) {
    switch(Action){
        case "InitializeClient":
            userId = parseInt(Value);
            break;
        case "StartTimer":
            startTimer(Value);
            break;
        case "StopTimer":
            stopTimer(Value);
            break;
        default:
            console.log("WebSocket got unhandled action: " + Action + " with Value:" + Value);
            break;
    }
}
