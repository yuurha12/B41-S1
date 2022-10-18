function ShowData() {
    let showName = document.getElementById('input-name').value;
    let showEmail = document.getElementById('input-email').value;
    let showPhone = document.getElementById('input-phone').value;
    let showSubject = document.getElementById('input-subject').value;
    let showMessage = document.getElementById('input-message').value;

    console.log(showName);
    console.log(showEmail);
    console.log(showPhone);
    console.log(showSubject);
    console.log(showMessage);

    if (showName == '') {
        return alert('Nama Harus diisi')
    }

    if (showEmail =='') {
        return alert('Email harus diisi')
    }

    if (showPhone =='') {
        return alert('Nomor telepon harus diisi')
    }

    if (showSubject == '') {
        return alert('Harus memilih salah satu')
    }

    if (showMessage == '') {
        return alert('Harus mengisi pesan')
    }

    let emailReceiver = 'hoki.ikki@gmail.com'

    let a = document.createElement('a');
    a.href = `mailto:${emailReceiver}?subject:${showSubject}&body= Hello, %0D%0A My name is ${showName}, ${showMessage}`

    a.click()


}