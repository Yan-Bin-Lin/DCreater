let reCAPTCHA_token = "6LcCuJwdAAAAAI2ITaq_01Xobie7X2FK9hTLuEvP";
// this is the id of the form
$("#SignForm").submit(function(e) {

    let url = $(this).attr('action');
    let username = $("#InputUserName").val();
    let password = $("#InputPassword").val();
    let data = new FormData();
    data.append('username', username);
    data.append('password', password);

    grecaptcha.ready(function() {
        grecaptcha.execute(reCAPTCHA_token, {action: 'submit'}).then(function(token) {
            // Add your logic to submit to your backend server here.
            data.append('response', token);
            $.ajax({
                method: "POST",
                url: url,
                data: data,
                processData: false,
                contentType: false,
                xhrFields: {
                    withCredentials: true
                },
                success: function(data){
                    console.log(data)
                    SetUserCookie(data.user)
                    window.location.reload();
                },
                error: function(jqXHR, textStatus, errorThrown){
                    console.log(jqXHR)
                    console.log(textStatus)
                    console.log(errorThrown)
                    $("#AlertWrongParam").show();
                }
            });
        });
    });

    e.preventDefault(); // avoid to execute the actual submit of the form.
});

// click on sign in
$("#SigninBtn").on("click", function(){
    $("#singContainer").show();
});

// click on x
$("#closeSignBtn").on("click", function(){
    HideSignContainer()
});

function HideSignContainer() {
    $("#singContainer").hide();
}

function obj_merge(obj1, obj2){
  Object.keys(obj2).forEach(function (k, index) {
    if (k in obj1){
      if (Array.isArray(obj1[k])){
        obj1[k].push(obj2[k]);
      } else {
        obj1[k] = [obj1[k], obj2[k]];
      }
    } else {
    	obj1[k] = [obj2[k]];
    }
	});
  return obj1
}

function SetUserCookie(user) {
    /*
    let sub = [["text", "subOid"], ["text", "subOuniquename"],["text", "subOnickname"],["text", "subOdescription"]];
    split_strs(user, sub)
    */
    user_data = {}
    for (let i = 0; i < user.length; i++) {
        obj_merge(user_data, user[i]);
    }
    document.cookie = "OwnerList=" + JSON.stringify(user_data) + ";Path=/;domain=.dcreater.com;Max-Age=2592000;SameSite=Lax";
}