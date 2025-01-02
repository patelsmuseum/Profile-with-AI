document.getElementById('button1').addEventListener('click', function() {
    fetch('/v1/welcome')
        .then(response => response.json())
        .then(data => alert(JSON.stringify(data)));
});

