package logger

type slog struct{}

func NewSlog(logPath string) (*slog, error) {
	return &slog{}, nil
}

func (s *slog) Error(msg string, args ...any) {}

func (s *slog) Info(msg string, args ...any) {
	print(msg)
}

func (s *slog) Debug(msg string, args ...any) {}

func (s *slog) Warn(msg string, args ...any) {}

func (s *slog) With(rgs ...any) Logger {
	return &slog{}
}
