function send() {
    var req = new XMLHttpRequest();

    req.open('GET', window.location.href+"request?url="+document.getElementById("url").value, true);

    req.addEventListener('load', function() {
        document.getElementById("response").innerHTML = req.responseText
    });

    req.send();
}
