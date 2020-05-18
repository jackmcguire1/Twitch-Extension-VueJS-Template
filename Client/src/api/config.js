function getGlobalConfig() {
  // eslint-disable-next-line no-undef
  var config = Twitch.ext.configuration;

  if (config.global) {
    var global = JSON.parse(config.global.content);
    return global;
  }

  return {};
}

function getDeveloperConfig() {
    // eslint-disable-next-line no-undef
  var config = Twitch.ext.configuration;

  if (config.broadcaster) {
    var developer = JSON.parse(config.developer.content);
    return developer;
  }

  return {};
}

function getBroadcasterConfig() {
  // eslint-disable-next-line no-undef
  var config = Twitch.ext.configuration;

  if (config.broadcaster) {
    var broadcaster = JSON.parse(config.broadcaster.content);
    return broadcaster;
  }

  return {};
}

export default {
  getBroadcasterConfig,
  getDeveloperConfig,
  getGlobalConfig,
};
