package fiberw

var DocString = `
<!DOCTYPE html>
<style type="text/css" media="screen">
   .code { 
   height: 100px
   }
</style>
<body>
   <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-GLhlTQ8iRABdZLl6O3oVMWSktQOp6b7In1Zl3/Jr59b6EGGoI1aFkw7cmDA6j6gD" crossorigin="anonymous">
   <nav class="navbar bg-body-tertiary">
      <div class="container-fluid">
         <span class="navbar-brand mb-0 h1">{{.ApplicationName}}</span>
      </div>
   </nav>
   <div class="container">
      <p class="lead mt-5">{{.Description}}</p>
      </br>
      <h3>API Routes</h3>
      {{ with .Groups }}
      {{range .}}	
      <div class="accordian card" id="main-{{.Name}}">
         <div class="accordion-item">
            <h2 class="accordion-header" id="{{.Name}}">
               <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#collapse-{{.Name}}" aria-expanded="true" >
                  <h3 class=".text-primary m-3">{{.Name}}</h3>
               </button>
            </h2>
            <div id="collapse-{{.Name}}" class="accordion-collapse collapse show" aria-labelledby="headingOne" data-bs-parent="#main-{{.Name}}">
               <div class="card-body">
                  <div class="accordion-body">
                     {{ with .Routes}}
                     {{range .}}	
                     <div class="accordion" id="accordionFlush{{.RouteType}}{{.GroupName}}{{.RouteName}}">
                        <div class="accordion-item">
                           <h2 class="accordion-header" id="flush-heading-{{.RouteType}}{{.GroupName}}{{.RouteName}}">
                              <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#flush-{{.RouteType}}{{.GroupName}}{{.RouteName}}" aria-expanded="false" aria-controls="flush-{{.RouteType}}{{.GroupName}}{{.RouteName}}">
                                 <b>
                                    <pre/><span class="badge text-bg-primary">{{.RouteType}}</span> {{.GroupName}}{{.RouteName}}</pre>
                                 </b>
                              </button>
                           </h2>
                           <div id="flush-{{.RouteType}}{{.GroupName}}{{.RouteName}}" class="accordion-collapse collapse"  data-bs-parent="#accordionFlush{{.RouteType}}{{.GroupName}}{{.RouteName}}">
                              <div class="accordion-body">
                                 <b>{{.Description}}</b>
                                 {{ $RouteType := .RouteType }}
                                 {{ $GroupName := .GroupName }}
                                 {{ $RouteName := .RouteName }}
                                 {{ if .HasParams }}
				 <p> <b> Path Parameters</b> </p>
                                 {{ with .Params }}
                                 {{ range $val := . }}
                                 <pre> {{ $val }} <input id="{{ $RouteType }}{{ $GroupName }}{{ $RouteName }}PARAM{{ $val }}" /> </pre>
                                 {{ end }}
                                 {{ end }}
                                 {{ end }}
                                 {{ if .HasQuery }}
				 <p> <b> Query Parameters </b> </p>
                                 {{ with .Queries }}
                                 {{ range $val := . }}
                                 <pre> {{ $val }} <input id="{{ $RouteType }}{{ $GroupName }}{{ $RouteName }}QUERY{{ $val }}" /> </pre>
                                 {{ end }}
                                 {{ end }}
                                 {{ end }}
                                 {{ if eq .RouteType "POST"}}
				 <p> <b> Request Body </b> </p>
                                 <div id="body-{{.RouteType}}{{.GroupName}}{{.RouteName}}" class="code">
                                    {{.Body}}
                                 </div>
                                 {{ end }}
                                 <br/>
                                 <button class="btn btn-secondary" id="btn-{{.RouteType}}{{.GroupName}}{{.RouteName}}">Send Request</button>
                                 <br/>
                                 <br/>
				 <p> <b> Response Body </b> </p>
                                 <div id="response-{{.RouteType}}{{.GroupName}}{{.RouteName}}" class="code">
                                 </div>
                              </div>
                           </div>
                        </div>
                     </div>
                     {{end}}
                     {{end}}
                  </div>
               </div>
            </div>
         </div>
         {{end}}
         {{end}}
      </div>
   </div>
   <script src="https://cdnjs.cloudflare.com/ajax/libs/ace/1.15.2/ace.min.js" integrity="sha512-9wsrxhzNVuW9XQgRmpSG9i23hheWGRZt0+M+T4vA/CXPLNEPCGsTXHaQi8/U5/gpuboqT0tFW+1hhUPzA4UHQA==" crossorigin="anonymous" referrerpolicy="no-referrer"></script>
   <script>
      // const baseUrl = "http://localhost:3000"
      const baseUrl = ""
      
      const makeGetRequest = async (url, responseId, queryParams, pathParams) => {
      	const editor = ace.edit(responseId, {
      		 mode: "ace/mode/json",
      	});
      	editor.setValue("Making API Call to " + url)
      	let query = {}
      	if(pathParams) {
      		for (const [k, v] of Object.entries(pathParams)) {
      			url = url.replace(":" + k, v)
      		}
      	}
      	if(queryParams) {
      		query = new URLSearchParams(queryParams)
      		url = url + "?" + query.toString()
      	}
      	const res = await fetch(url, {
      		method: "GET",
      	})
      	const json = await res.json();
      	console.log(json)
      	editor.setValue(JSON.stringify(json, null, 3))
      }
      
      const makePostRequest = async (url, body, responseId, queryParams, pathParams) => {
      	const editor = ace.edit(responseId, {
      		 mode: "ace/mode/json",
      	});
      	editor.setValue("Making API Call to " + url)
      	console.log(queryParams)
      	console.log("Params - ", pathParams)
      	if(pathParams) {
      		for (const [k, v] of Object.entries(pathParams)) {
      			url = url.replace(":" + k, v)
      		}
      	}
      	if(queryParams) {
      		query = new URLSearchParams(queryParams)
      		url = url + "?" + query.toString()
      	}
      	const res = await fetch(url, {
      		method: "POST",
      		body: body,
      		headers: {
      			"Content-Type": "application/json"
      		}
      	})
      	const json = await res.json();
      	console.log(json)
      	editor.setValue(JSON.stringify(json, null, 3))
      }
      
      for(const group of {{.Groups}}) {
      	for(const route of group.Routes) {
      		// Create a Body editor only the request type is POST (Change this for when adding headers?)
      		if(route.RouteType === "POST") { 
      			var editor = ace.edit("body-"+ route.RouteType + route.GroupName + route.RouteName, {
      				 mode: "ace/mode/json",
      			});

			editor.setValue(JSON.stringify(route.Body, null, 5), -1);
      		}
      		var responseEditor = ace.edit("response-" + route.RouteType + route.GroupName + route.RouteName, {
      			 mode: "ace/mode/json",
      		});
                  responseEditor.setValue(JSON.stringify(route.Returns, null, 5), -1)
      		responseEditor.setReadOnly(true)
      
      		document.getElementById('btn-' + route.RouteType + route.GroupName + route.RouteName).addEventListener('click', function(e) {
      			let query = false
      			if(route.HasQuery) {
      				query = {}
      				for (const q of route.Queries) {
      					const id = route.RouteType + route.GroupName + route.RouteName + "QUERY" + q
      					const value = document.getElementById(id)
      					query[q] = value.value
      					console.log("Setting query", query)
      				}
      			}
      			let params = false
      			if(route.HasParams) {
      				params = {}
      				for (const q of route.Params) {
      					const id = route.RouteType + route.GroupName + route.RouteName + "PARAM" + q
      					const value = document.getElementById(id)
      					params[q] = value.value
      				}
      			}
      			// Re-Init so we get the latest value when the function is called
      			const url = baseUrl + route.GroupName + route.RouteName
      			const responseId = "response-" + route.RouteType + route.GroupName + route.RouteName
      			console.log({query})
      			console.log({params})
      			if(route.RouteType === "GET") {
      				makeGetRequest(url, responseId, query, params)
      			}
      			else {
      				console.log(route)
      				var editor = ace.edit("body-" + route.RouteType + route.GroupName + route.RouteName, {
      					 mode: "ace/mode/json",
      				});
      				const postBody = editor.getValue()
      				console.log(postBody)
      				makePostRequest(url, postBody, responseId, query, params)
      			}
      		});
      	}
      }
      
   </script>
   <script></script>
   <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js" integrity="sha384-w76AqPfDkMBDXo30jS1Sgez6pr3x5MlQ1ZAGC+nuZB+EYdgRZgiwxhTBTkF7CXvN" crossorigin="anonymous"></script>
</body>
</html>
`
