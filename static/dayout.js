
(() => {
  'use strict'

  const forms = document.querySelectorAll('.needs-validation')

  
  Array.from(forms).forEach(form => {
    form.addEventListener('submit', event => {
      if (!form.checkValidity()) {
        event.preventDefault()
        event.stopPropagation()
      }

      form.classList.add('was-validated')
    }, false)
  })
})()









//  function sendEmail(){
//      Email.send({
//   Host : "smtp.gmail.com",
//   Username : "edwinsibypers@gmail.com",
//   Password : "109701EBECDA0FE2521157ABA36790B28781",
//   To : 'edwinsibyrajakumary@gmail.com',
//   From : document.getElementById("validationCustomUsername").value,
//   Subject : "New booking",
//   Body : "And this is the body"
//   }).then(
//   message => alert(message)
//   );
//   }