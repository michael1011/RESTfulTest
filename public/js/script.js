function registerListeners() {
    registerListener(document.getElementById("url"));
    registerListener(document.getElementById("body"));
    registerListener(document.getElementById("headers"));
}

function registerListener(element) {
    element.addEventListener("keydown", function (e) {
        if (e.keyCode === 13) {
            send()
        }
    });
}

var finished = false;

function send() {
    if (!finished) {
        finished = true;

        var loading = document.getElementById("loading");
        var resp = document.getElementById("response");

        resp.innerText = "";
        loading.style.visibility = 'visible';

        var req = new XMLHttpRequest();

        req.open('GET', window.location.href+"request?url="+document.getElementById("url").value+
            "&body="+document.getElementById("body").value+"&headers="+document.getElementById("headers").value, true);

        req.addEventListener('load', function() {
            loading.style.visibility = 'hidden';
            resp.innerText = req.responseText;

            $('pre code').each(function(i, block) {
                hljs.highlightBlock(block);
            });

            finished = false;
        });

        req.send();
    }

}
