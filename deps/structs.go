package deps

type (
	Dependency struct {
		Path  string
		Alias string
	}

	Echo struct {
		Dependency
		Middlewares Dependency
	}
	EchoRoutesView struct {
		Dependency
	}
	Net struct {
		Dependency
		Http Dependency
	}
	Fmt struct {
		Dependency
	}
	Std struct {
		Net
		Fmt
		Time   Dependency
		Maps   Dependency
		Slices Dependency
	}
	Gorm struct {
		Dependency
		Clause Dependency
	}
	I18n struct {
		Dependency
	}
	GookitValidate struct {
		Dependency
	}
	Project struct {
		Pkg
		Core
		Internal
		Cmd
	}
	Pkg struct {
		Dependency
		Config Dependency
	}
	Controllers struct {
		Dependency
		InitPages Dependency
	}
	Service struct {
		Dependency
		Controllers
		Translations Dependency
	}
	Cmd struct {
		Dependency
		Service
	}
	Domain struct {
		Dependency
		Models Dependency
	}
	Internal struct {
		Dependency
		Domain
	}
	Core struct {
		Dependency
		Converters  Dependency
		Validator   Dependency
		Logger      Dependency
		Connections Dependency
		CommonReq   Dependency
		AppErrors   Dependency
		CoreTpls    CoreTpls
	}
	CoreTpls struct {
		Dependency
		Toast Dependency
	}
	Container struct {
		Echo
		EchoRoutesView
		Std
		Gorm
		I18n
		GookitValidate
		Project
		MapStructure Dependency
		GoFakeIt     Dependency
		Uuid         Dependency
		Goption      Dependency
		GormPager    Dependency
		UDecimal     Dependency
	}
)
