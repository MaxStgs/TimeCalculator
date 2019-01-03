var nowTimer;


window.onload = function() {
    nowTimer = document.getElementById("current_time");
    updateNowTimer();
    setInterval(updateNowTimer, 1000);
};

function updateNowTimer() {
    nowTimer.textContent = new Date().toLocaleString();
}