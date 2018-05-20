if (window.console) {
    console.log("Some console stuff");
}

// Initialize Reveal

Reveal.initialize({
    slideNumber: true,
    transition: 'slide',
    viewDistance: 3
});

// Print pdf
var link = document.createElement( 'link' );
link.rel = 'stylesheet';
link.type = 'text/css';
link.href = window.location.search.match( /print-pdf/gi ) ? '/public/stylesheets/pdf.css' : '/public/stylesheets/paper.css';
document.getElementsByTagName( 'head' )[0].appendChild( link );
