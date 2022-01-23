//global
var ownerList

// rend for login user
let OwnerSelectNav = $("#OwnerSelectNav");
rend_owner();

// rend for owner list
function rend_owner() {
    ownerList = check_login();
    if (ownerList === undefined) {
        show_login();
        return;
    }

    for(let i = 0; i < ownerList.subOid.length; i++){
        OwnerSelectNav.append("<option value=\"" + ownerList.subOid[i] + "\">" + ownerList.subOnickname[i] + "</option>")
    }
    OwnerSelectNav.removeClass("hide");
    OwnerSelectNav.val(ownerList.subOid[0]);
    hide_login();
}

// return owner data if user has login, else return undefined
function check_login(){
    let ownerList = Cookies.get('OwnerList');
    if (ownerList !== undefined) {
        ownerList = JSON.parse(ownerList)
    }
    return ownerList;
}

// show login btn
function show_login(){
    $("#LoginNav").show();
}

// hide login btn
function hide_login(){
    $("#LoginNav").hide();
}

$("#OwnerHomePageBtn").click(function(e){
    window.location.href = "/blog/" + ownerList.subOuniquename[ownerList.subOid.indexOf(oid)];
})