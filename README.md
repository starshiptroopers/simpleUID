# uidgenerator

An another one random UID's generator with customized format, chars and length

Math.rand is using as a random generator, seed is initialized with time.Now().UnixNano() by default

## usage example:

        g := uidgenerator.New(&uidgenerator.Cfg{
            Alfa:      "1234567890",
            Format:    "XXX-XXXXXX-XXX",
            Validator: "[0-9]{3}-[0-9]{6}-[0-9]{3}",
        })
        uid1 := g.New()
        uid2 := g.New()

        uid3, err := g.Validate("111-222222-333")
      
## License

© Starship Troopers, 2020~time.Now

Released under the [MIT License] (https://github.com/starshiptroopers/uidgenerator/master/License)
