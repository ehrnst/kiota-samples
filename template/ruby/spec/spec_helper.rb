# frozen_string_literal: true

require 'microsoft_kiota_abstractions'
require_relative './../graphrubyv4/graphrubyv4.rb'
Dir["#{File.dirname(__FILE__)}/./../graphrubyv4/**/*.rb"].each { |f| load(f) }

RSpec.configure do |config|
  # Enable flags like --only-failures and --next-failure
  config.example_status_persistence_file_path = ".rspec_status"

  # Disable RSpec exposing methods globally on `Module` and `main`
  config.disable_monkey_patching!

  config.expect_with :rspec do |c|
    c.syntax = :expect
  end
end
