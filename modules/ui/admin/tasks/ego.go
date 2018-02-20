// Generated by ego.
// DO NOT EDIT

package tasks

import (
	"fmt"
	api "github.com/infinitbyte/gopa/core/http"
	"github.com/infinitbyte/gopa/core/i18n"
	"github.com/infinitbyte/gopa/core/model"
	"github.com/infinitbyte/gopa/modules/ui/common"
	"html"
	"io"
	"net/http"
)

var _ = fmt.Sprint("") // just so that we can keep the fmt import for now
func Index(w http.ResponseWriter, r *http.Request, domain string, status int, from, size, taskCount int, tasks []model.Task, domainsCount int, domains map[string]interface{}, kvs map[string]interface{}) error {
	_, _ = io.WriteString(w, "\n\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n\n")
	common.Head(w, "Tasks", "")
	_, _ = io.WriteString(w, "\n<link rel=\"stylesheet\" href=\"/static/assets/css/tasks.css\" />\n<script src=\"/static/assets/js/jquery.timeago.js\"></script>\n<script src=\"/static/assets/js/page/tasks.js\"></script>\n<script src=\"/static/assets/uikit-2.27.1/js/core/dropdown.min.js\"></script>\n<script src=\"/static/assets/uikit-2.27.1/js/components/pagination.min.js\"></script>\n\n")
	common.Body(w)
	_, _ = io.WriteString(w, "\n")
	common.Nav(w, r, "Tasks")
	_, _ = io.WriteString(w, "\n\n")

	paras := map[string]interface{}{}
	paras["host"] = domain
	paras["status"] = status

	_, _ = io.WriteString(w, "\n\n<div class=\"tm-middle\">\n\n    <div class=\"uk-container uk-container-center\">\n\n        <div class=\"uk-grid\" data-uk-grid-margin>\n\n            <div class=\"uk-width-2-10 uk-hidden-small\">\n                <div class=\"uk-alert\" ><span>Hosts</span></div>\n                <ul class=\"uk-list uk-list-striped\" style=\"max-height: 730px; overflow: scroll\">\n                    ")

	if domain != "" {

		_, _ = io.WriteString(w, "\n                    <li><a href=\"?\"><i class=\"uk-icon-filter\">:&nbsp;</i>")
		_, _ = fmt.Fprint(w, domain)
		_, _ = io.WriteString(w, "</a></li>\n                    ")

	} else {
		if len(domains) > 0 {
			for k, v := range domains {

				_, _ = io.WriteString(w, "\n                        ")
				_, _ = fmt.Fprint(w, GetDomainRow(k, v))
				_, _ = io.WriteString(w, "\n                        ")

			}
		}
	}

	_, _ = io.WriteString(w, "\n                </ul>\n            </div>\n\n            <div class=\"uk-width-8-10\">\n\n                <ul class=\"uk-tab\" data-uk-tab=\"\">\n                    ")

	if status > -1 {

		_, _ = io.WriteString(w, "\n                    <li><a href=\"/admin/tasks/?host=")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(domain)))
		_, _ = io.WriteString(w, "\"><i class=\"uk-icon-filter\">:&nbsp;</i>")
		_, _ = fmt.Fprint(w, model.GetTaskStatusText(status))
		_, _ = io.WriteString(w, "</a></li>\n                    ")

	} else {

		_, _ = io.WriteString(w, "\n                    <li ")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetActive(status, -1))))
		_, _ = io.WriteString(w, " ><a href=\"/admin/tasks/?host=")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(domain)))
		_, _ = io.WriteString(w, "\">All</a></li>\n                    <li  ")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetActive(status, 0))))
		_, _ = io.WriteString(w, "  aria-expanded=\"false\" class=\"\"><a href=\"/admin/tasks/?status=0&host=")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(domain)))
		_, _ = io.WriteString(w, "\">New(")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetStatusCount("0", kvs))))
		_, _ = io.WriteString(w, ")</a></li>\n                    <li  ")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetActive(status, 9))))
		_, _ = io.WriteString(w, "  aria-expanded=\"false\" class=\"\"><a href=\"/admin/tasks/?status=9&host=")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(domain)))
		_, _ = io.WriteString(w, "\">PendingFetch(")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetStatusCount("9", kvs))))
		_, _ = io.WriteString(w, ")</a></li>\n                    <li  ")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetActive(status, 3))))
		_, _ = io.WriteString(w, "  aria-expanded=\"false\" class=\"\"><a href=\"/admin/tasks/?status=3&host=")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(domain)))
		_, _ = io.WriteString(w, "\">Success(")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetStatusCount("3", kvs))))
		_, _ = io.WriteString(w, ")</a></li>\n                    <li data-uk-dropdown=\"{mode:'click'}\" aria-haspopup=\"true\" aria-expanded=\"false\">\n                        <a href=\"#\">More <i class=\"uk-icon-caret-down\"></i></a>\n                        <div class=\"uk-dropdown uk-dropdown-small\" aria-hidden=\"true\">\n                            <ul class=\"uk-nav uk-nav-dropdown\">\n                                <li  ")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetActive(status, 2))))
		_, _ = io.WriteString(w, "  aria-expanded=\"false\" class=\"\"><a href=\"/admin/tasks/?status=2&host=")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(domain)))
		_, _ = io.WriteString(w, "\">Failed(")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetStatusCount("2", kvs))))
		_, _ = io.WriteString(w, ")</a></li>\n                                <li  ")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetActive(status, 4))))
		_, _ = io.WriteString(w, "  aria-expanded=\"false\" class=\"\"><a href=\"/admin/tasks/?status=4&host=")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(domain)))
		_, _ = io.WriteString(w, "\">404(")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetStatusCount("4", kvs))))
		_, _ = io.WriteString(w, ")</a></li>\n                                <li  ")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetActive(status, 5))))
		_, _ = io.WriteString(w, "  aria-expanded=\"false\" class=\"\"><a href=\"/admin/tasks/?status=5&host=")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(domain)))
		_, _ = io.WriteString(w, "\">Redirected(")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetStatusCount("5", kvs))))
		_, _ = io.WriteString(w, ")</a></li>\n                                <li  ")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetActive(status, 6))))
		_, _ = io.WriteString(w, "  aria-expanded=\"false\" class=\"\"><a href=\"/admin/tasks/?status=6&host=")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(domain)))
		_, _ = io.WriteString(w, "\">Timeout(")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetStatusCount("6", kvs))))
		_, _ = io.WriteString(w, ")</a></li>\n                                <li  ")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetActive(status, 7))))
		_, _ = io.WriteString(w, "  aria-expanded=\"false\" class=\"\"><a href=\"/admin/tasks/?status=7&host=")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(domain)))
		_, _ = io.WriteString(w, "\">Duplicated(")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetStatusCount("7", kvs))))
		_, _ = io.WriteString(w, ")</a></li>\n                                <li  ")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetActive(status, 8))))
		_, _ = io.WriteString(w, "  aria-expanded=\"false\" class=\"\"><a href=\"/admin/tasks/?status=8&host=")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(domain)))
		_, _ = io.WriteString(w, "\">Interrupted(")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(GetStatusCount("8", kvs))))
		_, _ = io.WriteString(w, ")</a></li>\n                            </ul>\n                        </div>\n                    </li>\n                    <li class=\"uk-tab-responsive uk-active uk-hidden\" aria-haspopup=\"true\" aria-expanded=\"false\"><a>Item</a><div class=\"uk-dropdown uk-dropdown-small\" aria-hidden=\"true\"><ul class=\"uk-nav uk-nav-dropdown\"></ul><div></div></div></li>\n\n                    ")

	}

	_, _ = io.WriteString(w, "\n                    <script language=\"javascript\">\n                        function locate(url){\n                            if(url!=undefined)\n                            {\n                                window.location.replace(url)\n                            }\n                        }\n                        $(function () {\n                            $('.uk-tab').on('change.uk.tab', function(e, active, previous) {\n                                locate(active.context.firstChild.href);\n                            });\n                            $('.uk-nav-dropdown').click(function(e){\n                                locate(e.target.href);\n                            });\n                        });\n                    </script>\n\n                </ul>\n\n                <div class=\"uk-overflow-container\">\n                    <table id=\"tasks\" class=\"uk-table uk-table-hover uk-table-striped\" cellspacing=\"0\" width=\"100%\">\n                    <thead>\n                    <tr>\n                        <th>URL</th>\n                        <th>LastUpdate</th>\n                        <th>NextCheck</th>\n                        <th>Status</th>\n                    </tr>\n                    </thead>\n                    <tbody id=\"records\">\n                    ")
	if len(tasks) > 0 {
		for _, task := range tasks {

			_, _ = io.WriteString(w, "\n                    ")
			_, _ = fmt.Fprint(w, GetTaskRow(task))
			_, _ = io.WriteString(w, "\n                    ")

		}
	}

	_, _ = io.WriteString(w, "\n                    </tbody>\n                </table>\n\n                    ")
	_, _ = fmt.Fprint(w, api.GetPagination(from, size, taskCount, "", paras))
	_, _ = io.WriteString(w, "\n\n                </div>\n            </div>\n        </div>\n\n    </div>\n\n\n\n</div>\n\n<!-- modal start -->\n\n<div id=\"create-task-modal\" class=\"uk-modal\" aria-hidden=\"true\" style=\"display: none; overflow-y: auto;\">\n    <div class=\"uk-modal-dialog uk-modal-dialog-blank\">\n        <button class=\"uk-modal-close uk-close\" type=\"button\"></button>\n        <div class=\"uk-grid uk-flex-middle\" data-uk-grid-margin=\"\">\n            <div class=\"uk-width-medium-1-3 uk-height-viewport uk-cover-background uk-row-first\" style=\"background:#000000;\"></div>\n            <div class=\"uk-width-medium-2-3 uk-row-first \">\n                <h1>Create Task</h1>\n                <div class=\"uk-width-medium-3-4\">\n                    <form class=\"uk-form\">\n                        <fieldset data-uk-margin>\n                            <textarea type=\"text\" placeholder=\"Paste your urls here, one line one url\" rows=\"10\" class=\"uk-width-1-1\"></textarea>\n                            <input type=\"text\" placeholder=\"MaxDepth, eg: 30\">\n                            <input type=\"text\" placeholder=\"MaxBreadth, eg: 3\">\n\n                            <input type=\"checkbox\" id=\"form-s-c\" />\n                            <label for=\"form-s-c\">Follow Sub-domain</label>\n                            <br/>\n                            <input name=\"include-ext\" type=\"checkbox\" id=\"img-ext\" />\n                            <label for=\"img-ext\">Include Images</label>\n                            <input  name=\"include-ext\" type=\"checkbox\" id=\"css-ext\" />\n                            <label for=\"css-ext\">Include Css</label>\n                            <input  name=\"include-ext\" type=\"checkbox\" id=\"js-ext\" />\n                            <label for=\"js-ext\">Include JS</label>\n                        </fieldset>\n                        <button class=\"uk-button\">Create</button>\n                    </form>\n                </div>\n            </div>\n        </div>\n    </div>\n</div>\n\n<!-- modal end -->\n\n")
	common.Footer(w)
	_, _ = io.WriteString(w, "\n")
	return nil
}
func View(w http.ResponseWriter, r *http.Request, task model.Task) error {
	_, _ = io.WriteString(w, "\n\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n\n")
	common.Head(w, "Task", "")
	_, _ = io.WriteString(w, "\n<link rel=\"stylesheet\" href=\"/static/assets/css/tasks.css\" />\n<script src=\"/static/assets/js/jquery.timeago.js\"></script>\n<script src=\"/static/assets/js/page/tasks.js\"></script>\n<script src=\"/static/assets/uikit-2.27.1/js/components/pagination.min.js\"></script>\n\n")
	common.Body(w)
	_, _ = io.WriteString(w, "\n")
	common.Nav(w, r, "Tasks")
	_, _ = io.WriteString(w, "\n<style>\n     strong{\n        margin-right: 10px;;\n    }\n</style>\n<div class=\"tm-middle\">\n\n    <div class=\"uk-container uk-container-center\">\n\n        <div class=\"uk-grid\" data-uk-grid-margin=\"\">\n\n            <div class=\"tm-main uk-width-medium-3-4\">\n                <div class=\"uk-overflow-container\">\n                    <div class=\"tm-middle-header\">\n                        <div class=\"uk-grid \" data-uk-grid-margin=\"\">\n                            <div class=\"uk-width-1-1\">\n                                <h2>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Task"))))
	_, _ = io.WriteString(w, "</h2>\n                            </div>\n                        </div>\n                    </div>\n\n                    <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("ID"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.ID)))
	_, _ = io.WriteString(w, "</div>\n                    <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Host"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.Host)))
	_, _ = io.WriteString(w, "</div>\n                    <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Url"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.Url)))
	_, _ = io.WriteString(w, "</div>\n                    <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("OriginalUrl"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.OriginalUrl)))
	_, _ = io.WriteString(w, "</div>\n                    <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Reference"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.Reference)))
	_, _ = io.WriteString(w, "</div>\n                    <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Status"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(model.GetTaskStatusText(task.Status))))
	_, _ = io.WriteString(w, "</div>\n                    <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Depth"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.Depth)))
	_, _ = io.WriteString(w, "</div>\n                    <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Breadth"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.Breadth)))
	_, _ = io.WriteString(w, "</div>\n                    <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Created"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.Created)))
	_, _ = io.WriteString(w, "</div>\n                    <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Updated"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.Updated)))
	_, _ = io.WriteString(w, "</div>\n                    <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("LastFetch"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.LastFetch)))
	_, _ = io.WriteString(w, "</div>\n                    <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("LastCheck"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.LastCheck)))
	_, _ = io.WriteString(w, "</div>\n                    <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("NextCheck"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.NextCheck)))
	_, _ = io.WriteString(w, "</div>\n\n\n                    <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("PipelineConfigID"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.PipelineConfigID)))
	_, _ = io.WriteString(w, "</div>\n\n                    <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Message"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.Message)))
	_, _ = io.WriteString(w, "</div>\n\n                </div>\n\n                <hr />\n\n                <div>\n                    <div class=\"tm-middle-header\">\n                        <div class=\"uk-grid \" data-uk-grid-margin=\"\">\n                            <div class=\"uk-width-1-1\">\n                                <h2>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Snapshot"))))
	_, _ = io.WriteString(w, "(")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.SnapshotCount)))
	_, _ = io.WriteString(w, ")</h2>\n                            </div>\n                        </div>\n                        <div class=\"uk-grid grid-margin\">\n                            <div>\n                                <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("SnapshotVersion"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.SnapshotVersion)))
	_, _ = io.WriteString(w, "</div>\n                                <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("SnapshotID"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.SnapshotID)))
	_, _ = io.WriteString(w, "</div>\n                                <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("SnapshotHash"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.SnapshotHash)))
	_, _ = io.WriteString(w, "</div>\n                                <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("SnapshotSimHash"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.SnapshotSimHash)))
	_, _ = io.WriteString(w, "</div>\n                                <div class=\"p\"><strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("SnapshotCreated"))))
	_, _ = io.WriteString(w, ":</strong>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.SnapshotCreated)))
	_, _ = io.WriteString(w, "</div>\n                            </div>\n                        </div>\n                    </div>\n                    <div class=\"tm-middle-body\">\n                        <div class=\"uk-overflow-container\">\n                            <table class=\"uk-table uk-table-hover border\">\n                                <thead>\n                                <tr>\n                                    <th class=\"uk-text-center\">")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("#"))))
	_, _ = io.WriteString(w, "</th>\n                                    <th class=\"uk-text-center\">")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Created"))))
	_, _ = io.WriteString(w, "</th>\n                                    <th class=\"uk-text-center\">")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Size"))))
	_, _ = io.WriteString(w, "</th>\n                                    <th class=\"uk-text-center\">")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Version"))))
	_, _ = io.WriteString(w, "</th>\n                                    <th class=\"uk-text-center\">")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Hash"))))
	_, _ = io.WriteString(w, "</th>\n                                    <th class=\"uk-text-center\">")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Title"))))
	_, _ = io.WriteString(w, "</th>\n                                    <th class=\"uk-text-center\">")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Action"))))
	_, _ = io.WriteString(w, "</th>\n                                </tr>\n                                </thead>\n                                <tbody>\n                                ")
	for i, v := range task.Snapshots {
		_, _ = io.WriteString(w, "\n                                <tr class=\"uk-text-center\">\n                                    <td>")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i)))
		_, _ = io.WriteString(w, "</td>\n                                    <td>")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(v.Created)))
		_, _ = io.WriteString(w, "</td>\n                                    <td>")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(v.Size)))
		_, _ = io.WriteString(w, "</td>\n                                    <td>")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(v.Version)))
		_, _ = io.WriteString(w, "</td>\n                                    <td>")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(v.Hash)))
		_, _ = io.WriteString(w, "</td>\n                                    <td>")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(v.Title)))
		_, _ = io.WriteString(w, "</td>\n                                    <td><a class=\"uk-button uk-button-small uk-button-primary\" target=\"_blank\" href=\"/snapshot/")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(v.ID)))
		_, _ = io.WriteString(w, "\">View</a></td>\n                                </tr>\n                                ")
	}
	_, _ = io.WriteString(w, "\n                                </tbody>\n                            </table>\n                        </div>\n                    </div>\n                </div>\n            </div>\n            <div class=\"tm-sidebar uk-width-medium-1-4 uk-hidden-small uk-row-first\">\n                ")
	if task.LastScreenshotID != "" {
		_, _ = io.WriteString(w, "\n                    <ul class=\"tm-nav uk-nav\" data-uk-nav=\"\">\n                        <h2>")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(i18n.GetLocalizedText("Screenshot"))))
		_, _ = io.WriteString(w, "</h2>\n                    </ul>\n                    <img style=\"max-width: 100%\" src=\"/screenshot/")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(task.LastScreenshotID)))
		_, _ = io.WriteString(w, "\" />\n                ")
	}
	_, _ = io.WriteString(w, "\n            </div>\n        </div>\n        </div>\n\n</div>\n\n")
	common.Footer(w)
	_, _ = io.WriteString(w, "\n")
	return nil
}
