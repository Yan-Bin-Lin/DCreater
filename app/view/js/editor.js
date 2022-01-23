let editor;
// option of blog type
let blogType;
// initial value of blog content
let initialContent = "";
// check hashtag
HandleNewPost();

function HandleNewPost() {
    // editflag that create editor
    let editFlag = ["#NewBlog", "#UpdateBlog"];
    if (editFlag.includes(window.location.hash)) {
        HandleShowForm();
        editor = ShowEditor();
    } else {
        HandleHideForm();
        HideEditor();
    }

}

// create a new editor
function CreateEditor() {
    const {Editor} = toastui;
    const {codeSyntaxHighlight, colorSyntax, tableMergedCell} = Editor.plugin;

    return Editor.factory({
        el: document.querySelector('#editor'),
        usageStatistics: false,
        height: '500px',
        initialEditType: 'wysiwyg',
        previewStyle: 'tab',
        initialValue: initialContent,
        plugins: [codeSyntaxHighlight, colorSyntax, tableMergedCell],
        hooks: {
            'addImageBlobHook': imageUpload
        }
    })
}

// handle for editor upload image
function imageUpload(blob, callback) {
    let data = new FormData();
    let path = window.location.pathname.split("/").slice(1);
    let oid = path.length > 1 ? meta.boid : meta.oid;
    let name = geTimeNowStr();
    let url = "https://dcreater.com/file/img/" + oid + "/" + name;
    data.append('content', blob);
    data.append('fileName', name);
    data.append('oid', oid);
    $.ajax({
        url: url,
        data: data,
        cache: false,
        contentType: false,
        processData: false,
        method: 'POST',
        mimeType: 'multipart/form-data',
        xhrFields: {
            withCredentials: true
        },
        success: function (data) {
            // run callback
            callback(url);
        },
        error: function (jqXHR, textStatus, errorThrown) {
            alert("sorry, something error in upload image. Please try again later");
        }
    });
    function geTimeNowStr() {
        let date = new Date(Date.now());
        return date.toISOString();
    }
}

// prevent color picker reload page
$("#NewBlogForm").submit(function (e) {
    e.preventDefault();
});

// create a new blog or update exist blog
$("#EditorSubmitBtn").on("click", function (e) {
    let blogType = {"1": "project", "2": "article"};

    let name = $("#NewBlogName").val();
    let newName = (name === meta.bname ? "" : name);
    let description = $("#NewBlogDescription").val();
    let type = blogType[$("#TypeSelect").val()];
    let content = new Blob([editor.getMarkdown()], {type: 'text/plain'});
    let url = window.location.pathname + (window.location.hash === "#UpdateBlog" ? "" : "/" + name);
    let path = decodeURI(window.location.pathname).split("/").slice(1);
    let oid = path.length > 2 ? meta.boid : meta.oid;
    let method = 'POST';
    let newSuperPath =  GetMidURL(path)
    let superid = (window.location.hash === "#UpdateBlog" ? meta.bsuper : (path.length > 2 ? meta.bid : 0));

    let data = new FormData();
    data.append('descript', description);
    data.append('blogType', type);
    data.append("superid", superid);
    data.append("oid", oid);
    if (type === "article") {
        data.append("content", content);
    }
    if (window.location.hash === "#UpdateBlog") {
        // update
        data.append("bid", meta.bid);
        data.append("newsuperid", -1);
        data.append("newsuperUrl", newSuperPath);
        data.append("newname", newName);
        method = 'put';
    }

    $.ajax({
        url: url,
        data: data,
        cache: false,
        contentType: false,
        processData: false,
        method:  method,
        mimeType: 'multipart/form-data',
        withCredentials: true,
        success: function (data) {
            // redirect
            window.location.href = window.location.hash === "#UpdateBlog" ? ["", "blog", newSuperPath, (newName === "" ? meta.bname : newName)].join("/") : url;
        },
        error: function (jqXHR, textStatus, errorThrown) {
            alert("sorry, something error in create blog. Please try again later");
        }
    });

    e.preventDefault();
});

/**
 * @return {string}
 */
function GetMidURL(path){
    path = path.slice(1, path.length - 1);
    path.unshift("");
    return path.join("/");
}

// choose type event
$('#TypeSelect').on('click', function () {
    blogType = $(this).val();
});

//  hashtag event
$(window).on('hashchange', function () {
    HandleNewPost();
});

function HandleShowForm() {
    HideBlogList();
    ShowReturnBtn();
    //HideNewBlogBtn();
    $("#editorContainer").show();
    FillEditor();
}

function HandleHideForm() {
    ShowBlogList();
    HideReturnBtn();
    ShowNewBlogBtn();
    HideForm();
}

// show blog or project form
function ShowForm() {
    $("#editorContainer").show();
    FillEditor();
}

// fill form of update blog
function FillEditor() {
    let content = $('#content');
    if (content.length > 0) {
        $("#NewBlogName").val(meta.bname);
        $("#NewBlogDescription").val(meta.bdescription);
        $("#TypeSelect").val(meta.btype).prop('disabled', true);
        initialContent = $("#metaContent").html();
    }
}

// Hide blog or project form
function HideForm() {
    $("#editorContainer").hide();
}

function HideEditor() {
    $('#editor').children().remove();
}

function ShowEditor() {
    return CreateEditor();
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

function ShowNewBlogBtn() {
    $("#NewBlogNav").show();
    /*
    if (meta.btype === 1) {
        $("#NewBlogNav").show();
    } else {
        $("#NewBlogNav").hide();
    }
    */
    $("#UpdateBlogNav").show();
}

function HideNewBlogBtn() {
    $("#NewBlogNav").hide();
    $("#UpdateBlogNav").hide();
}