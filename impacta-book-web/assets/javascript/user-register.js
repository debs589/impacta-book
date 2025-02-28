$('#register-form').on('submit', createUser)

function createUser(event) {
    event.preventDefault();
    console.log("Testing");

    if ($('#password').val() != $('#confirme-password').val()) {
        alert("The passwords don't match!");
        return;
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            nickname: $('#nickname').val(),
            password: $('#password').val()
        }
    });

}