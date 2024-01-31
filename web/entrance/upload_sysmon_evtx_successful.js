// Show success message
Swal.fire({
    icon: "success",
    title: "EVTX file uploaded successfully!",
    text: "Redirecting to welcome page...",
    showConfirmButton: false,
    timer: 2000,
    timerProgressBar: true
}).then(() => {
    // Redirect to welcome.html
    window.location.href = "/entrance";
});
