extends PanelContainer

enum DAY{
	LUN,
	MAR,
	MIE,
	JUE,
	VIE,
	SAB
}

const DAYNAME = [
	"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"
]

const DAYCODE = {
	Monday = 0,
	Tuesday = 1,
	Wednesday = 2,
	Thursday = 3,
	Friday = 4,
	Saturday = 5,
}

# estructura de una hora (array)
# 0 day -> DAY
# 1 start -> 1-24
# 2 end -> 1-24

class Course:
	var ctg: String
	var level: String
	var name: String
	var nrc: String
	var hours: Array = []

# create courses

var courses: Array = []
var activeCourses: Array = []
var categories: Array = []
var categoryControls: Array = []

var listitem = preload("res://controls/item.tscn")

onready var courseContainer: VBoxContainer = $"%courseContainer"
onready var selectedCourseContainer: TextEdit = $"%selectedCourseContainer"
onready var categorySelection: OptionButton = $"%categorySelection"

#colors
var GRAY: Color = Color(.99,.99,.99,1)
export var GRAY_LIGHT: Color = Color(0.9,0.9,0.9,1)
var WHITE: Color = Color(1,1,1,1)
var BLACK: Color = Color(0,0,0,1)
var RED: Color = Color(1,0,0,1)
var REDB: Color = Color(1,0,0,1)

export var DOTS_COLOR: Color = Color(1,0,0,1)
export var DAY_COLORS: Array = [Color(1,0,0,1)]

# start, its 0 because its relative to its origin
var xstr:float = 40
var ystr:float = 40

# size
var bigwd:float = rect_size.x
var bight:float = rect_size.y

var wd:float = bigwd - xstr -4
var ht:float = bight - ystr -4

#font
export var fontTags:DynamicFont
export var fontCourses:DynamicFont

# Called when the node enters the scene tree for the first time.
func _ready():
	
	# generate controls
	genControls()

func _draw():
	
	bigwd = rect_size.x
	bight = rect_size.y

	wd = bigwd - xstr -4
	ht = bight - ystr -4
	
	var xmarks:float = 6 # 5 days
	var ymarks:float = 22 - 4 # start at 4, end at 22
	
	# clear all labels
	var volconlab: Node2D = $VolatileLabel
	free_children(volconlab as Node)
	
	# dynamic label
	var mylabel:Label
	
	# whole background
	draw_rect(Rect2(0,0, bigwd, bight), GRAY_LIGHT, true, 2)
	
	# background
	draw_rect(Rect2(xstr,ystr, wd,ht), WHITE, true, 2)
	
	# vertical grid and days
	
	for i in range(xmarks):
		var x1 = xstr + wd/xmarks * i
		var texty = ystr - 25
		
		# h line
		draw_line(Vector2(x1, ystr), Vector2(x1, ystr + ht), GRAY_LIGHT, 1)
		
		# text
		mylabel = Label.new()
		mylabel.rect_position = Vector2(x1 + wd/xmarks/2, texty)
		mylabel.set_align(mylabel.ALIGN_CENTER)
		mylabel.set_valign(mylabel.VALIGN_BOTTOM)
		mylabel.grow_horizontal = Control.GROW_DIRECTION_BOTH
		mylabel.grow_vertical = Control.GROW_DIRECTION_END
		
		mylabel.text = DAYNAME[i]
		mylabel.set("custom_fonts/font", fontTags)
		mylabel.set("custom_colors/font_color", BLACK)
		volconlab.add_child(mylabel)
		
	# horizontal grid and hours
	
	for i in range(ymarks +1):
		var y1 = ystr + ht/ymarks * i
		var textx = xstr - 10
		var hour = (i + 4) % 13 + floor((i + 4) / 13)
		
		# h line
		draw_line(Vector2(xstr, y1), Vector2(xstr + wd, y1), GRAY_LIGHT, 1)
		
		# text
		mylabel = Label.new()
		mylabel.rect_position = Vector2(textx, y1)
		mylabel.set_align(mylabel.ALIGN_RIGHT)
		mylabel.set_valign(mylabel.VALIGN_CENTER)
		mylabel.grow_horizontal = Control.GROW_DIRECTION_BEGIN
		mylabel.grow_vertical = Control.GROW_DIRECTION_BOTH
		
		mylabel.text = str(hour) 
		mylabel.set("custom_fonts/font", fontTags)
		mylabel.set("custom_colors/font_color", BLACK)
		volconlab.add_child(mylabel)
	
	# draw courses
	
	for i in range(courses.size()):
		
		if (!activeCourses[i]):
			continue
		
		var c: Course = courses[i]
		var w = wd/xmarks
		var h = ht/ymarks 
		var color = DAY_COLORS[randi() % DAY_COLORS.size()]
		
		# estructura de una hora (array)
		# 0 day -> DAY
		# 1 start -> 4-22
		# 2 end -> 4-22

		for j in range(c.hours.size()):
			
			var hour: Array = c.hours[j]
			
			# day
			var x1 = xstr + w * DAYCODE[hour[0]]
			# hora
			var y1 = ystr + ht * ((hour[1] -4.0) / ymarks)
			var yend = ystr + ht * ((hour[2] -4.0) / ymarks)
		
			# box
			draw_rect(Rect2(x1, y1, w, yend - y1), color, true, 2)
			draw_rect(Rect2(x1, y1, w, yend - y1), BLACK, false, 1)
			
			# text
			mylabel = Label.new()
			mylabel.rect_position = Vector2(x1, y1)
			mylabel.set_align(mylabel.ALIGN_LEFT)
			mylabel.set_valign(mylabel.VALIGN_BOTTOM)
			mylabel.grow_horizontal = Control.GROW_DIRECTION_END
			mylabel.grow_vertical = Control.GROW_DIRECTION_END
			mylabel.margin_right = x1+w
			mylabel.autowrap = true
			
			mylabel.text = c.name + " " + c.nrc
			mylabel.set("custom_fonts/font", fontCourses)
			mylabel.set("custom_colors/font_color", BLACK)
			volconlab.add_child(mylabel)

func genControls():
	
	# clean
	
	free_children(courseContainer as Node)
	
	# fill categories
	
	#categorySelection
	
	var c  # child
	var cb # checkbox
	var ctg
	
	# instance for every course
	for i in range(courses.size()):
		
		ctg = courses[i].ctg
		
		
		# build categories
		if !categories.has(ctg):
			categories.append(ctg)
			categorySelection.add_item(ctg)
		
		if categories[categorySelection.selected] != ctg:
			continue

		c = listitem.instance()
		c.get_node("HBoxContainer/Label").text = courses[i].level + " " + courses[i].nrc + " " + courses[i].name
		
		# connect signal
		cb = c.get_node("HBoxContainer/CheckBox")
		cb.connect("toggled", self, "checkboxToggleSignal", [i])
		
		
		# add controls to categories
		
		
		courseContainer.add_child(c)
		
func checkboxToggleSignal(buttonPressed, courseId: int):
	activeCourses[courseId] = buttonPressed
	
	# force redraw
	self.update()
	self.showSelectedCourses()
	
func showSelectedCourses():
	
	selectedCourseContainer.text = ""
	
	for i in range(courses.size()):
		
		if (activeCourses[i]):
			var c = courses[i]
			selectedCourseContainer.text += c.nrc + " " + c.name + "\n"
	
	
func create_class(ctg: String, name: String, nrc: String, hours: Array, level: String) -> Course:
	var c = Course.new()
	c.ctg   = ctg
	c.name  = name
	c.nrc   = nrc
	c.hours = hours
	c.level = level
	
	activeCourses.append(false)
	return c

func parse_json_to_courses(json: Array):
	
	if (!json.size()):
		return
		
	courses = []
	
	for c in json:
		
		var hours: Array
		
		# convert hours
		for h in c["hour"]:
			hours.append([h["day"], h["start"], h["end"]])
		
		courses.append(create_class(c["ctg"], c["title"], c["nrc"], hours, str(c["level"])))
		
	# reset selected courses
	for i in range(activeCourses.size()):
		activeCourses[i] = false
		
	genControls()
	self.update()
	
func free_children(node: Node):
	
	var children = node.get_children()
	for i in range(children.size()):
		var c = children[i]
		node.remove_child(c)
		c.queue_free()
