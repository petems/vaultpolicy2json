Given(/^I have "([^"]*)" command installed$/) do |command|
  is_present = system("which #{ command} > /dev/null 2>&1")
  raise "Command #{command} is not present in the system" if not is_present
end

Given("a build of {string}") do |string|

end
