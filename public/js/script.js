function send() {
    var loading = document.getElementById("loading");
    var resp = document.getElementById("response");

    resp.innerText = "";
    loading.style.visibility = 'visible';

    var req = new XMLHttpRequest();

    req.open('GET', window.location.href+"request?url="+document.getElementById("url").value, true);

    req.addEventListener('load', function() {
        loading.style.visibility = 'hidden';
        resp.innerText = req.responseText
    });

    req.send();
}
