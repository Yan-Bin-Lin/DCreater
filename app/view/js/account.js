// global
let oid, ownername, nickname, descryption

if(ownerList !== undefined){
    SetNowOwner(ownerList.subOid[0]);
    OwnerSelectNav.removeClass("hide");
}

// check hashtag
HandleNewOwner();

// owner list on change
OwnerSelectNav.on('change', function() {
    SetNowOwner(OwnerSelectNav.val());
    FillForm();
});

function SetNowOwner(in_id) {
    oid = in_id;
    let i = ownerList.subOid.indexOf(oid);
    ownername = ownerList.subOuniquename[i];
    nickname = ownerList.subOnickname[i];
    descryption = i < ownerList.subOdescription.length ? ownerList.subOdescription[i] : "";
}

//  hashtag event
$(window).on('hashchange', function () {
    HandleNewOwner();
});

function HandleNewOwner() {
    // editflag that create editor
    let editFlag = ["#NewOwner", "#UpdateOwner"];

    if (editFlag.includes(window.location.hash)) {
        HandleShowOwnerForm();
    } else {
        HandleHideOwnerForm();
    }

}

$("#DeleteOwnerBtn").on("click", function() {
    if (confirm("確定刪除?") !== true ){ 
        return; 
    }
    let oid = OwnerSelectNav.val();

    let data = new FormData();
    data.append("oid", oid);

    $.ajax({
        url: "/blog/" + ownerList.subOuniquename[ownerList.subOid.indexOf(oid)],
        method: 'DELETE',
        data: data,
        cache: false,
        contentType: false,
        processData: false,
        withCredentials: true,
        success: function(result) {
            $("option[value=\"" + oid + "\"]").remove();
            alert("刪除成功");
        },
        error: function(jqXHR, textStatus, errorThrown){
            console.log(jqXHR);
            console.log(textStatus);
            console.log(errorThrown);
            alert("刪除失敗，請稍後再試");
        },
    });
});

// create a new owner or update exist owner
$("#OwnerSubmitBtn").on("click", function (e) {
    let method = "post";
    let in_name = $("#NewOwnerName").val();

    let data = new FormData();
    data.append("nickname", $("#NewNickName").val());
    data.append("descript", $("#NewOwnerDescription").val());
    data.append("oid", oid);

    if (window.location.hash === "#UpdateOwner") {
        // update
        data.append("newuniname", (in_name === ownername ? "" : in_name));
        method = 'put';
    }

    $.ajax({
        url: "/blog/" + (window.location.hash === "#NewOwner" ? in_name : ownername),
        data: data,
        cache: false,
        contentType: false,
        processData: false,
        method:  method,
        withCredentials: true,
        success: function (data) {
            // redirect
            window.location.href = "https://dcreater.com";
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log(jqXHR);
            console.log(textStatus);
            console.log(errorThrown);
            alert("sorry, something error. Please try again later");
        }
    });

    e.preventDefault();
});

function HandleShowOwnerForm() {
    HideBlogList();
    ShowReturnBtn();
    HideNewOwnerBtn();
    ShowForm();
    FillForm();
}

function HandleHideOwnerForm() {
    ShowBlogList();
    HideReturnBtn();
    ShowNewOwnerBtn();
    HideForm();
}

function FillForm() {
    if(window.location.hash === "#NewOwner") {
        return;
    }
    $("#NewOwnerName").val(ownername);
    $("#NewNickName").val(nickname);
    $("#NewOwnerDescription").val(descryption);
}

// show Owner or project form
function ShowForm() {
    $("#formContainer").show();
}

// Hide Owner or project form
function HideForm() {
    $("#formContainer").hide();
}

function HideBlogList() {
    $("main.container").hide();
}

function ShowBlogList() {
    $("main.container").show();
}

function ShowReturnBtn() {
    $("#ReturnNav").show();
}

function HideReturnBtn() {
    $("#ReturnNav").hide();
}

function ShowNewOwnerBtn() {
    $("#NewOwnerNav").show();
    $("#UpdateOwnerNav").show();
}

function HideNewOwnerBtn() {
    $("#NewOwnerNav").hide();
    $("#UpdateOwnerNav").hide();
}