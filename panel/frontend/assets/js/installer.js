//TODO: move this to css https://stackoverflow.com/questions/8294400/css-animations-with-delay-for-each-child-element

document.getElementById('start-install-btn').addEventListener("click", () => {
    document.getElementById('start-container').classList.add('fadeOut')
    setTimeout(() => {
        document.getElementById('start-container').style.display = 'none';
        document.getElementById('settings-container').classList.add('fadeIn')
        document.getElementById('settings-container').classList.remove('d-none');
    
    }, 1500);
});


Promise.all(
    document.getElementById('installer-description').getAnimations().map(
      function(animation) {
        return animation.finished
      }
    )
  ).then(
    function() {
      return document.getElementById('start-install-btn').classList.add('slideLeft');
    }
  );