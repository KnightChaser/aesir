document.getElementById('upload_button').addEventListener('click', function () {
    // Get the file input element and name input element
    var fileInput = document.getElementById('formFile');
    var nameInput = document.getElementById('floatingInput');
    
    // Check if either file or name is not filled
    if (fileInput.files.length === 0 || nameInput.value.trim() === '') {
        Swal.fire({
            icon: 'error',
            title: 'Incomplete Form',
            text: 'Please upload the file and designate the name.',
        });
        // Prevent form submission
        return false;
    }

    // Check if a file is selected
    if (fileInput.files.length > 0) {
        var fileName = fileInput.files[0].name;
        // Check if the file has a valid extension
        if (!fileName.endsWith('.evtx')) {
            Swal.fire({
                icon: 'error',
                title: 'Unsupported File Extension',
                text: 'Please upload a file with the ".evtx" extension.',
            });
            // Prevent form submission
            return false;
        }
        
        // Ask the user to confirm file upload
        Swal.fire({
            title: 'Confirm Upload',
            text: 'Are you sure you want to upload the file?',
            icon: 'question',
            showCancelButton: true,
            confirmButtonText: 'Yes, upload it!',
            cancelButtonText: 'Cancel'
        }).then((result) => {
            if (result.isConfirmed) {
            
                // Show progress to the user with a toast
                const Toast = Swal.mixin({
                    toast: true,
                    position: 'top-end',
                    iconColor: 'white',
                    showConfirmButton: false,
                    timer: 2000,
                    timerProgressBar: true
                });

                (async () => {
                    await Toast.Fire({
                        icon: 'info',
                        title: 'Uploading file...'
                    })
                })()

                // Proceed with form submission
                document.getElementById('uploadForm').submit();
                
            }
        });
    }
});
