package branch


import (
	"net/http"
	"strings"
	// readme was indirectly imported
	// readme has already been imported in the main.go file as a package 
	"readme"
)

// this page handles the ABOUT BRANCH when creating a new branch
// What the page gives are three LINKS.
// 1. for adding the about branch which is a textarea
// 2. for guides on how to add about branch, as it contains more special features to integrate on the about page
// 3. for previewing the about branch file before creating a new branch


func AboutBranchLink (r *http.Request, name string ) string {

	collect := []string{}

	// parent
	collect = append(collect, "<div class='about_brnch_addjss9890'><div class='about_brnch_addjss9890-chld'>")

	// link
	collect = append(collect, `<div class='about_brnch_addjss9891'>`)

	if name == "about" {
		collect = append(collect, `
		    <div class='about_brnch_lnk090'>
			    <div class='about_brnch_lnk090-chld'>
				    <div class='aboutbtnjss870 about_brnch_lnk091 aboutbtnjss870 about_brnch_lnk091_active' data_label="about">
					    <span style="pointer-events:none;">About</span>
					</div>
				</div> 
			</div>
		`)
	}else{
		collect = append(collect, `
		    <div class='about_brnch_lnk090'>
			    <div class='about_brnch_lnk090-chld'>
				    <div class='aboutbtnjss870 about_brnch_lnk091' data_label="about">
					    <span style="pointer-events:none;">About</span>
					</div>
				</div> 
			</div>
		`)
	}


	if name == "guide" {
		collect = append(collect, `
		    <div class='about_brnch_lnk090'>
			    <div class='about_brnch_lnk090-chld'>
				    <div class='aboutbtnjss870 about_brnch_lnk091 aboutbtnjss870 about_brnch_lnk091_active' data_label="guide">
					    <span style="pointer-events:none;">Guides</span>
					</div>
				</div> 
			</div>
		`)
	}else{
		collect = append(collect, `
		    <div class='about_brnch_lnk090'>
			    <div class='about_brnch_lnk090-chld'>
				    <div class='aboutbtnjss870 about_brnch_lnk091' data_label="guide">
					    <span style="pointer-events:none;">Guides</span>
					</div>
				</div> 
			</div>
		`)
	}

	if name == "preview" {
		collect = append(collect, `
		    <div class='about_brnch_lnk090'>
			    <div class='about_brnch_lnk090-chld'>
				    <div class='aboutbtnjss870 about_brnch_lnk091 aboutbtnjss870 about_brnch_lnk091_active' data_label="preview">
					    <span style="pointer-events:none;">Preview</span>
					</div>
				</div> 
			</div>
		`)
	}else{
		collect = append(collect, `
		    <div class='about_brnch_lnk090'>
			    <div class='about_brnch_lnk090-chld'>
				    <div class='aboutbtnjss870 about_brnch_lnk091' data_label="preview">
					    <span style="pointer-events:none;">Preview</span>
					</div>
				</div> 
			</div>
		`)
	}

	collect = append(collect, "</div>")
	// link

	if name == "about" {
		collect = append(collect, `
		    <div class="about_brnch_addjss9892">
		        <div class="about_brnch_addjss9891_details">
		            <textarea name='about' id='abutres90' placeholder='Description'>`+strings.TrimSpace(r.FormValue("about"))+`</textarea> 
					<div id="txterr890"></div>
				</div>
		   </div>
		`)
	}
	
	if name == "guide" {
		collect = append(collect, `
		    <div class="about_brnch_addjss9892">
		        <div class="about_brnch_addjss9891_details">
		            <h2>Guide for this shit</h2>
				</div>
		   </div>
		`)
	}


	if name == "preview" {
		prev := strings.TrimSpace(r.FormValue("preview"));

		if len(prev) > 0 {
			collect = append(collect, `
				<div class="about_brnch_addjss9892">
					<div style="background:#f2f2f2;padding:5px;border-radius:4px;' class="about_brnch_addjss9891_details">
						<div>`+readme.MoldReadMe(false, prev)+`</div>
					</div>
			   </div>
			`)
		}else{
			collect = append(collect, `
				<div class="about_brnch_addjss9892">
					<div class="about_brnch_addjss9891_details">
						<h2>Add text to preview</h2>
					</div>
			   </div>
			`)
		}

	}

	collect = append(collect, "</div></div>")
	// parent


	return strings.Join(collect, "")
}