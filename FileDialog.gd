extends FileDialog

export var mainControllerPath: NodePath
var mainController: PanelContainer

onready var acceptDialog: AcceptDialog = get_node("AcceptDialog3")

func _ready():
	mainController = get_node(mainControllerPath)

func _on_Button_pressed():
	self.popup_centered(Vector2(800,600))

func _on_FileDialog_file_selected(path):
	
	# load file
	
	var f = File.new()
	
	if !f.file_exists(path):
		acceptDialog.dialog_text =  "File path does not exist"
		acceptDialog.popup_centered(Vector2(300,100))
		return

	f.open(path, File.READ)
	var filetext = f.get_as_text()
	var json = JSON.parse(filetext)
	
	# check file
	
	if json.error != OK:
		push_error("Unexpected results.")
		
		acceptDialog.dialog_text = "Error in line " + str(json.error_line) + ": " + json.error_string 
		acceptDialog.popup_centered(Vector2(300,100))
		return
		
	elif typeof(json.result) != TYPE_ARRAY:
		acceptDialog.dialog_text = "Invalid JSON"
		acceptDialog.popup_centered(Vector2(300,100))
		return
	
	print(json.result)
	mainController.parse_json_to_courses(json.result)
