package Builder

type Robot struct {
	HasHead bool
	HasArm  bool
	HasFoot bool
}

func NewRobot(hasHead bool, hasArm bool, hasFoot bool) *Robot {
	return &Robot{
		HasHead: hasHead,
		HasArm:  hasArm,
		HasFoot: hasFoot,
	}
}

type RobotBuilder struct {
	robot *Robot
}

func (b *RobotBuilder) BuildHead() *RobotBuilder {
	b.robot.HasHead = true
	return b
}
func (b *RobotBuilder) BuildArm() *RobotBuilder {
	b.robot.HasArm = true
	return b
}
func (b *RobotBuilder) BuildFoot() *RobotBuilder {
	b.robot.HasFoot = true
	return b
}
func (b *RobotBuilder) Return() *Robot {
	return b.robot
}
