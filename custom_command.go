package redis

func SetFirstKeyPos(cmd Cmder, pos int8) {
	cmd.setFirstKeyPos(pos)
}