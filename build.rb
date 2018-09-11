arch_targets = ["386","amd64"]
os_targets = ["darwin", "linux", "windows"]

version = "01"
program_name = 'creochain'
file_to_compile = './main/main.go'

os_targets.each do |os|
	arch_targets.each do |arch|

		file_name = program_name +'_'+ version + '_' + os + '_' + arch
		ENV['GOOS'] = os
		ENV['GOARCH'] = arch

		command = 'go build -o build/' + file_name + " " + file_to_compile

		system(command)

		if os == 'windows'
			File.rename('./build/'+file_name, './build/'+file_name+'.exe')
    end
  end
end