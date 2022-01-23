let blogTypeInvert = {
    1: "project",
    2: "article"
};

let filler = {
    "blog-list-owner": ["text", "onickname"],
    "owner-href": ["href", "ouniquename"],
};

// global
var meta = JSON.parse($("#meta").text());
let path = window.location.pathname.split("/").slice(1);
if (window.location.pathname === "/blog") {
    rend_root();
} else {
    rend_blog();
    rend_content();
}

function rend_root() {
    let sub = {
        "blog-list-title": ["text", "bname"],
        "blog-list-description": ["text", "bdescription"],
        "blog-list-createtime": ["text", "bcreatetime"],
        "blog-href": ["href", "burlpath"],
    };

    if (meta !== null){
        let result = "";
        for (let i = 0; i < meta.length; i++) {
            result += fill_content("#blog-list-tmpl", meta[i], Object.assign(filler, sub), 1);
        }
        $("div.blog-list").html(result);
    }
}

function rend_blog() {
    let sub = {
        "blog-list-title": ["text", "subBname"],
        "blog-list-description": ["text", "subBdescription"],
        "blog-list-createtime": ["text", "subBcreatetime"],
        "blog-href": ["href", "subBname"],
    };

    split_strs(meta, Object.values(sub));
    $("div.blog-list").html(fill_content("#blog-list-tmpl", meta, Object.assign(filler, sub), meta.subBname.length));
}

// split string to array
function split_strs(data, targets) {
    targets.forEach(target => {
        if (!Array.isArray(data[target[1]])) {
            data[target[1]] = data[target[1]].split("  ");
        }
    })
}

// get template and fill with data
function fill_content(tmplid, data, filler, len) {
    let tmpl = get_tmpl(tmplid);
    if (tmpl.length === 0) {
        return ""
    }
    let result = "";
    for (let i = 0; i < len; i++) {
        fill_tmpl(tmpl, data, filler, i);
        result += tmpl[0].outerHTML;
    }
    return result
}

function get_tmpl(target) {
    let html = $(target).html();
    return $(html).clone();
}

function fill_tmpl(tmpl, data, filler, index) {
    let value = "";
    for (let [k, v] of Object.entries(filler)) {

        if (Array.isArray(data[v[1]])) {
            value = data[v[1]][index];
        } else {
            value = data[v[1]];
        }

        if (v[0] === "text") {
            fill_text(tmpl.find("." + k), value);
        } else {
            fill_attr(tmpl.find("." + k), v[0], value);
        }
    }
}

function fill_text(node, value) {
    node.text(value);
}

function fill_attr(node, attr, value) {
    if (attr === "href") {
        fill_href(node, value);
    } else {
        node.attr(attr, value);
    }
}

function fill_href(node, value) {
    console.log(node)
    console.log(value)
    if (
        (window.location.pathname !== "/blog" && node.attr("class") === "blog-href") || 
        (window.location.pathname === "/blog" && node.attr("class") === "owner-href blog-list-owner")) {
            value = window.location.pathname + "/" + value;
    }
    else {
        value = window.location.pathname + value;
    }
    node.attr("href", value);
}

// rend for blog content
function rend_content() {
    let content = $('#content');
    if (content.length <= 0) {
        return;
    }

    let context = content.html();

    const {Editor} = toastui;
    const {codeSyntaxHighlight, tableMergedCell} = Editor.plugin;

    // create viewer
    const viewer = Editor.factory({
        el: document.querySelector('#content'),
        viewer: true,
        height: '500px',
        initialValue: context,
        plugins: [codeSyntaxHighlight, tableMergedCell]
    });

    // hide origin
    content.after("<div id='metaContent' style='display:none'>" + context + "</div>");
}

// return btn event
$("#DeleteBlogBtn").on("click", function() {
    if (confirm("確定刪除?") !== true ){ 
        return; 
    }
    // check url
    let path = window.location.pathname.split('/');
    if(path.length < 2 || path[3] === ""){
        return;
    }

    let data = new FormData();
    data.append("oid", OwnerSelectNav.val());
    data.append("bid", meta.bid);

    $.ajax({
        url: window.location.pathname,
        type: 'DELETE',
        data: data,
        cache: false,
        contentType: false,
        processData: false,
        withCredentials: true,
        success: function(result) {
            window.location.href = "https://dcreater.com" + path.slice(0, path.length - 1).join('/');
        },
        error: function(jqXHR, textStatus, errorThrown){
            console.log(jqXHR);
            console.log(textStatus);
            console.log(errorThrown);
        },
    });
});

/*
// return btn event
$("#ReturnNav").on("click", function() {
    removeHash();
});

function removeHash () {
    var scrollV, scrollH, loc = window.location;
    if ("pushState" in history)
        history.pushState("", document.title, loc.pathname + loc.search);
    else {
        // Prevent scrolling by storing the page's current scroll offset
        scrollV = document.body.scrollTop;
        scrollH = document.body.scrollLeft;

        loc.hash = "";

        // Restore the scroll offset, should be flicker free
        document.body.scrollTop = scrollV;
        document.body.scrollLeft = scrollH;
    }
}*/